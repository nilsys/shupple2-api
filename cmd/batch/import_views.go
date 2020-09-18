package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/pkg/errors"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"golang.org/x/oauth2"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/analytics/v3"

	"github.com/urfave/cli"
)

const (
	gaMetrics = "ga:pageviews"
	// MEMO: 特に意味はない
	gaStartDate     = "2015-01-01"
	gaEndDate       = "today"
	gaDimensions    = "ga:pagePath"
	gaSort          = "-ga:pageviews"
	gaMaxResult     = 99999
	gaSamplingLevel = "HIGHER_PRECISION"
	gaDateFmt       = "2006-01-02"
	tourismPrefix   = "/tourism"
	postsLimit      = 100
	ssmKey          = "sw-media-api-analytics-config"
	spanWeekly      = "weekly"
	spanMonthly     = "monthly"
	flagNameMedia   = "media"
	flagNameSpan    = "span"
)

type (
	Row struct {
		Path  string
		Views int
	}

	Entry struct {
		ID    int
		Type  model.EntityType
		Slug  string
		Score int
	}
)

var (
	pageRe = regexp.MustCompile(`(Page|ページ|Part) (\d)`)
)

func (b *Batch) cliImportViews() cli.Command {
	return cli.Command{
		Name:  "import_views",
		Usage: "指定したmediaのviewsをgoogle analyticsからimport/updateする",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     flagNameMedia,
				Usage:    "取り込むviewsのmediaを指定(Post,Review,Vlog,Feature,All(全て))",
				Required: true,
			},
			cli.StringFlag{
				Name:     flagNameSpan,
				Usage:    "取り込む日付範囲(weekly,monthly,all)",
				Required: true,
			},
		},
		Action: b.importViews,
	}
}

func (b *Batch) importViews(c *cli.Context) error {
	// ssmからjsonを取得する
	sess := session.Must(session.NewSession())
	svc := ssm.New(
		sess,
		aws.NewConfig().WithRegion(b.Config.AWS.Region),
	)
	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(ssmKey),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return errors.Wrap(err, "failed to get from parameter store")
	}

	key := []byte(*res.Parameter.Value)

	// ssmで取得した文字列->[]byteからjwtconfigをgen
	jwtConf, err := google.JWTConfigFromJSON(
		key,
		analytics.AnalyticsReadonlyScope,
	)
	if err != nil {
		return errors.Wrap(err, "failed to gen jwt config")
	}

	// TODO: deprecated
	httpClient := jwtConf.Client(oauth2.NoContext)
	gasvc, err := analytics.New(httpClient)
	if err != nil {
		return errors.Wrap(err, "failed to gen google analytics service")
	}

	mediaType, err := model.ParseEntityType(c.String(flagNameMedia))
	if err != nil {
		return errors.Wrap(err, "can't parse entity type")
	}
	span := c.String(flagNameSpan)
	gaToday := time.Now().Format(gaDateFmt)
	var gaStart string
	var gaEnd string

	switch span {
	case spanWeekly:
		gaStart = time.Now().Add(-168 * time.Hour).Format(gaDateFmt)
		gaEnd = gaToday
	case spanMonthly:
		gaStart = time.Now().Add(-720 * time.Hour).Format(gaDateFmt)
		gaEnd = gaToday
	default:
		gaStart = gaStartDate
		gaEnd = gaEndDate
	}

	totalResults := 1
	currentTotal := 0
	var results []*analytics.GaData

	for currentTotal < totalResults {
		gares, err := gasvc.Data.Ga.Get("ga:"+strconv.Itoa(b.Config.GoogleAnalytics.ViewID), gaStart, gaEnd, gaMetrics).Dimensions(gaDimensions).Sort(gaSort).StartIndex(int64(currentTotal + 1)).MaxResults(gaMaxResult).SamplingLevel(gaSamplingLevel).Do()
		if err != nil {
			return errors.Wrap(err, "failed to get analytics data")
		}

		totalResults = int(gares.TotalResults)
		results = append(results, gares)
		currentTotal += len(gares.Rows)
	}

	return b.aggregate(mediaType, span, results)
}

// TODO: 一旦2の固定値
func possibilityPostPaths(post *entity.Post) [2]string {
	a := tourismPrefix + "/" + post.Slug
	b := "/" + post.Slug
	return [2]string{a, b}
}

func possibilityVlogPath(vlog *entity.Vlog) [2]string {
	a := tourismPrefix + "/movie/" + strconv.Itoa(vlog.ID)
	b := "/movie/" + strconv.Itoa(vlog.ID)
	return [2]string{a, b}
}

func possibilityFeaturePath(feature *entity.Feature) [2]string {
	a := tourismPrefix + "/feature/" + strconv.Itoa(feature.ID)
	b := "/feature/" + strconv.Itoa(feature.ID)
	return [2]string{a, b}
}

func possibilityReviewPath(review *entity.Review) string {
	return "/review/" + strconv.Itoa(review.ID)
}

// 対象となるパスのみ抽出
// MEMO: ここでパスから/tourism削除しても良いかも？
func analyticsDataToRows(data *analytics.GaData) []Row {
	rows := make([]Row, 0)

	for _, row := range data.Rows {

		// 2ページ目以降のパスは含めない
		page := pageRe.FindAllStringSubmatch(row[0], 1)
		if len(page) >= 1 {
			continue
		}

		views, _ := strconv.Atoi(row[1])

		rows = append(rows, Row{
			Path:  row[0],
			Views: views,
		})
	}
	return rows
}

func aggregateReview(review *entity.Review, rows []Row) *Row {
	result := &Row{}
	reviewPath := possibilityReviewPath(review)

	for _, row := range rows {
		if strings.HasSuffix(row.Path, reviewPath) {
			if result.Path == "" {
				result.Path = row.Path
			}
			result.AddViews(row.Views)
		}
	}

	if len(result.Path) == 0 {
		return nil
	}

	return result
}

func aggregate(rows []Row, paths [2]string) *Row {
	result := &Row{}
	for _, row := range rows {
		if strings.HasPrefix(row.Path, paths[0]) || strings.HasPrefix(row.Path, paths[1]) {
			if result.Path == "" {
				result.Path = row.Path
			}
			result.AddViews(row.Views)
		}
	}

	if len(result.Path) == 0 {
		return nil
	}

	return result
}

func (r *Row) AddViews(views int) {
	r.Views += views
}

func (b *Batch) aggregate(mediaType model.EntityType, span string, garesResults []*analytics.GaData) error {
	rows := make([]Row, 0)
	for _, gares := range garesResults {
		rows = append(rows, analyticsDataToRows(gares)...)
	}

	switch mediaType {
	case model.EntityTypePost:
		entries, err := b.postAggregate(rows, span)
		if err != nil {
			return errors.Wrap(err, "failed post aggregate")
		}
		return b.adjustmentViews(entries, span)
	case model.EntityTypeReview:
		return b.reviewAggregate(rows, span)
	case model.EntityTypeVlog:
		entries, err := b.vlogAggregate(rows, span)
		if err != nil {
			return errors.Wrap(err, "failed vlog aggregate")
		}
		return b.adjustmentViews(entries, span)
	case model.EntityTypeFeature:
		entries, err := b.featureAggregate(rows, span)
		if err != nil {
			return errors.Wrap(err, "failed feature aggregate")
		}
		return b.adjustmentViews(entries, span)
	default:
		entries := make([]Entry, len(rows))
		postEntries, err := b.postAggregate(rows, span)
		if err != nil {
			return errors.Wrap(err, "failed post aggregate")
		}
		err = b.reviewAggregate(rows, span)
		if err != nil {
			return errors.Wrap(err, "failed post aggregate")
		}
		vlogEntries, err := b.vlogAggregate(rows, span)
		if err != nil {
			return errors.Wrap(err, "failed vlog aggregate")
		}
		featureEntries, err := b.featureAggregate(rows, span)
		if err != nil {
			return errors.Wrap(err, "failed feature aggregate")
		}

		entries = append(entries, postEntries...)
		entries = append(entries, vlogEntries...)
		entries = append(entries, featureEntries...)

		return b.adjustmentViews(entries, span)
	}
}

func (b *Batch) adjustmentViews(entries []Entry, span string) error {
	sort.Slice(entries, func(i, j int) bool {
		return utf8.RuneCountInString(entries[i].Slug) > utf8.RuneCountInString(entries[j].Slug)
	})
	for _, entry := range entries {
		slugScore := 0
		for _, e := range entries {
			if entry.Slug != e.Slug && strings.HasPrefix(e.Slug, entry.Slug) {
				slugScore += e.Score
			}
		}
		entry.Score -= slugScore

		switch span {
		case spanWeekly:
			switch entry.Type {
			case model.EntityTypePost:
				if err := b.PostCommandRepository.UpdateWeeklyViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update review.weekly_views")
				}
			case model.EntityTypeReview:
				if err := b.ReviewCommandRepository.UpdateWeeklyViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update review.weekly_views")
				}
			case model.EntityTypeVlog:
				if err := b.VlogCommandRepository.UpdateWeeklyViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update vlog.weekly_views")
				}
			case model.EntityTypeFeature:
				if err := b.FeatureCommandRepository.UpdateWeeklyViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update feature.weekly_views")
				}
			}
		case spanMonthly:
			switch entry.Type {
			case model.EntityTypePost:
				if err := b.PostCommandRepository.UpdateMonthlyViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update review.monthly_views")
				}
			case model.EntityTypeReview:
				if err := b.ReviewCommandRepository.UpdateMonthlyViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update review.monthly_views")
				}
			case model.EntityTypeVlog:
				if err := b.VlogCommandRepository.UpdateMonthlyViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update vlog.monthly_views")
				}
			case model.EntityTypeFeature:
				if err := b.FeatureCommandRepository.UpdateMonthlyViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update feature.monthly_views")
				}
			}
		default:
			switch entry.Type {
			case model.EntityTypePost:
				if err := b.PostCommandRepository.UpdateViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update review.views")
				}
			case model.EntityTypeReview:
				if err := b.ReviewCommandRepository.UpdateViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update review.views")
				}
			case model.EntityTypeVlog:
				if err := b.VlogCommandRepository.UpdateViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update vlog.views")
				}
			case model.EntityTypeFeature:
				if err := b.FeatureCommandRepository.UpdateViewsByID(entry.ID, entry.Score); err != nil {
					return errors.Wrap(err, "failed to update feature.views")
				}
			}
		}
	}
	return nil
}

func (b *Batch) postAggregate(rows []Row, span string) ([]Entry, error) {
	entries := make([]Entry, 0)
	lastID := 0
	for {
		posts, err := b.PostQueryRepository.FindByLastID(lastID, postsLimit)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find by lastID")
		}
		if len(posts) == 0 {
			break
		}
		for _, post := range posts {
			row := aggregate(rows, possibilityPostPaths(post))
			if row == nil {
				continue
			}
			switch span {
			case spanWeekly:
				if err := b.PostCommandRepository.UpdateWeeklyViewsByID(post.ID, row.Views); err != nil {
					return nil, errors.Wrap(err, "failed to update weekly_views")
				}
			case spanMonthly:
				if err := b.PostCommandRepository.UpdateMonthlyViewsByID(post.ID, row.Views); err != nil {
					return nil, errors.Wrap(err, "failed to update monthly_views")
				}
			default:
				if err := b.PostCommandRepository.UpdateViewsByID(post.ID, row.Views); err != nil {
					return nil, errors.Wrap(err, "failed to update views")
				}
			}
			entries = append(entries, Entry{
				ID:    post.ID,
				Type:  model.EntityTypePost,
				Slug:  post.Slug,
				Score: row.Views,
			})
		}
		lastID = posts[len(posts)-1].ID
	}
	return entries, nil
}

func (b *Batch) reviewAggregate(rows []Row, span string) error {
	// TODO: lastID対応
	reviews, err := b.ReviewQueryRepository.FindAll()
	if err != nil {
		return errors.Wrap(err, "failed to find all review")
	}
	for _, review := range reviews {
		row := aggregateReview(review, rows)
		if row == nil {
			continue
		}
		switch span {
		case spanWeekly:
			if err := b.ReviewCommandRepository.UpdateWeeklyViewsByID(review.ID, row.Views); err != nil {
				return errors.Wrap(err, "failed to update review.weekly_views")
			}
		case spanMonthly:
			if err := b.ReviewCommandRepository.UpdateMonthlyViewsByID(review.ID, row.Views); err != nil {
				return errors.Wrap(err, "failed to update review.monthly_views")
			}
		default:
			if err := b.ReviewCommandRepository.UpdateViewsByID(review.ID, row.Views); err != nil {
				return errors.Wrap(err, "failed to update review.views")
			}
		}
	}
	return nil
}

func (b *Batch) vlogAggregate(rows []Row, span string) ([]Entry, error) {
	// TODO: lastID対応
	entries := make([]Entry, 0)
	vlogs, err := b.VlogQueryRepository.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to find all vlog")
	}
	for _, vlog := range vlogs {
		row := aggregate(rows, possibilityVlogPath(vlog))
		if row == nil {
			continue
		}
		switch span {
		case spanWeekly:
			if err := b.VlogCommandRepository.UpdateWeeklyViewsByID(vlog.ID, row.Views); err != nil {
				return nil, errors.Wrap(err, "failed to update vlog.weekly_views")
			}
		case spanMonthly:
			if err := b.VlogCommandRepository.UpdateMonthlyViewsByID(vlog.ID, row.Views); err != nil {
				return nil, errors.Wrap(err, "failed to update vlog.monthly_views")
			}
		default:
			if err := b.VlogCommandRepository.UpdateViewsByID(vlog.ID, row.Views); err != nil {
				return nil, errors.Wrap(err, "failed to update vlog.views")
			}
		}
		entries = append(entries, Entry{
			ID:    vlog.ID,
			Type:  model.EntityTypeVlog,
			Slug:  vlog.Slug,
			Score: row.Views,
		})
	}
	return entries, nil
}

func (b *Batch) featureAggregate(rows []Row, span string) ([]Entry, error) {
	// TODO: lastID対応
	entries := make([]Entry, 0)
	features, err := b.FeatureQueryRepository.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to find all feature")
	}
	for _, feature := range features {
		row := aggregate(rows, possibilityFeaturePath(feature))
		if row == nil {
			continue
		}
		switch span {
		case spanWeekly:
			if err := b.FeatureCommandRepository.UpdateWeeklyViewsByID(feature.ID, row.Views); err != nil {
				return nil, errors.Wrap(err, "failed to update feature.weekly_views")
			}
		case spanMonthly:
			if err := b.FeatureCommandRepository.UpdateMonthlyViewsByID(feature.ID, row.Views); err != nil {
				return nil, errors.Wrap(err, "failed to update feature.monthly_views")
			}
		default:
			if err := b.FeatureCommandRepository.UpdateViewsByID(feature.ID, row.Views); err != nil {
				return nil, errors.Wrap(err, "failed to update feature.views")
			}
		}
		entries = append(entries, Entry{
			ID:    feature.ID,
			Type:  model.EntityTypeFeature,
			Slug:  feature.Slug,
			Score: row.Views,
		})
	}
	return entries, nil
}
