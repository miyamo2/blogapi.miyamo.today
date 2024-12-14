package dto

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"net/url"
)

var (
	ErrInvalidateArticlesInDto = errors.New("invalidate articles in dto")
	ErrInvalidateTagsInDto     = errors.New("invalidate tags in dto")
)

type ArticleInDto struct {
	id string
}

// Id returns id.
func (i ArticleInDto) Id() string {
	return i.id
}

// IsInDto is a marker for in dto.
func (i ArticleInDto) IsInDto() {}

// NewArticleInDto constructor of ArticleInDto.
func NewArticleInDto(id string) ArticleInDto {
	return ArticleInDto{
		id: id,
	}
}

// ArticlesInDto is a dto for articles.
type ArticlesInDto struct {
	first  int
	last   int
	after  string
	before string
}

func (i ArticlesInDto) First() int {
	return i.first
}

func (i ArticlesInDto) Last() int {
	return i.last
}

func (i ArticlesInDto) After() string {
	return i.after
}

func (i ArticlesInDto) Before() string {
	return i.before
}

type ArticlesInDtoOption func(*ArticlesInDto) error

// ArticlesInWithFirst specifies how many articles to retrieve from the beginning.
func ArticlesInWithFirst(first int) ArticlesInDtoOption {
	return func(d *ArticlesInDto) error {
		if d.last != 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if last is set, first cannot be set.")
		}
		if d.before != "" {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if before is set, first cannot be set.")
		}
		d.first = first
		return nil
	}
}

// ArticlesInWithLast specifies how many articles to retrieve from the end.
func ArticlesInWithLast(last int) ArticlesInDtoOption {
	return func(d *ArticlesInDto) error {
		if d.first != 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if first is set, last cannot be set.")
		}
		if d.after != "" {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if after is set, last cannot be set.")
		}
		d.last = last
		return nil
	}
}

// ArticlesInWithAfter specifies which cursor to get as the starting point.
//
// NOTE: must always be executed after ArticlesInWithFirst
// NOTE: if ArticlesInWithLast or ArticlesInWithBefore was executed, it will be returned ErrInvalidateArticlesInDto.
func ArticlesInWithAfter(after string) ArticlesInDtoOption {
	return func(d *ArticlesInDto) error {
		if d.last != 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if last is set, after cannot be set.")
		}
		if d.before != "" {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if before is set, after cannot be set.")
		}
		if d.first == 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if first is not set, after cannot be set.")
		}
		d.after = after
		return nil
	}
}

// ArticlesInWithBefore specifies which cursor to get as the starting point.
//
// NOTE: must always be executed after ArticlesInWithLast
// NOTE: if ArticlesInWithFirst or ArticlesInWithAfter was executed, it will be returned ErrInvalidateArticlesInDto.
func ArticlesInWithBefore(before string) ArticlesInDtoOption {
	return func(d *ArticlesInDto) error {
		if d.first != 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if first is set, before cannot be set.")
		}
		if d.after != "" {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if after is set, before cannot be set.")
		}
		if d.last == 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDto, "if last is not set, before cannot be set.")
		}
		d.before = before
		return nil
	}
}

// NewArticlesInDto constructor of ArticlesInDto.
// if options are invalid, return ErrInvalidateArticlesInDto.
func NewArticlesInDto(options ...ArticlesInDtoOption) (ArticlesInDto, error) {
	d := ArticlesInDto{}
	for _, option := range options {
		err := option(&d)
		if err != nil {
			return ArticlesInDto{}, err
		}
	}
	return d, nil
}

type Article struct {
	id           string
	title        string
	body         string
	thumbnailUrl url.URL
	createdAt    synchro.Time[tz.UTC]
	updatedAt    synchro.Time[tz.UTC]
}

// IsOutDto is a marker for out dto.
func (a Article) IsOutDto() {}

// Id returns id.
func (a Article) Id() string {
	return a.id
}

// Title returns title.
func (a Article) Title() string {
	return a.title
}

// ThumbnailUrl returns thumbnail url.
func (a Article) ThumbnailUrl() url.URL {
	return a.thumbnailUrl
}

// CreatedAt returns created at.
func (a Article) CreatedAt() synchro.Time[tz.UTC] {
	return a.createdAt
}

// UpdatedAt returns updated at.
func (a Article) UpdatedAt() synchro.Time[tz.UTC] {
	return a.updatedAt
}

func NewArticle(id, title, body string, thumbnailUrl url.URL, createdAt, updatedAt synchro.Time[tz.UTC]) Article {
	return Article{
		id:           id,
		title:        title,
		body:         body,
		thumbnailUrl: thumbnailUrl,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

type Tag struct {
	id   string
	name string
}

// IsOutDto is a marker for out dto.
func (t Tag) IsOutDto() {}

func (t Tag) Id() string {
	return t.id
}

func (t Tag) Name() string {
	return t.name
}

func NewTag(id, name string) Tag {
	return Tag{
		id:   id,
		name: name,
	}
}

type ArticleTag struct {
	Article
	tags []Tag
}

// Body returns body.
func (a ArticleTag) Body() string {
	return a.body
}

// Tags returns tags.
func (a ArticleTag) Tags() []Tag {
	return a.tags
}

func NewArticleTag(id, title, body string, thumbnailUrl url.URL, createdAt, updatedAt synchro.Time[tz.UTC], tags []Tag) ArticleTag {
	return ArticleTag{
		Article: NewArticle(id, title, body, thumbnailUrl, createdAt, updatedAt),
		tags:    tags,
	}
}

type ArticleOutDto struct {
	article ArticleTag
}

func NewArticleOutDto(article ArticleTag) ArticleOutDto {
	return ArticleOutDto{
		article: article,
	}
}

func (o ArticleOutDto) Article() ArticleTag {
	return o.article
}

// IsOutDto is a marker for out dto.
func (o ArticleOutDto) IsOutDto() {}

// ArticlesOutDto is a dto for articles.
type ArticlesOutDto struct {
	articles   []ArticleTag
	byBackward bool
	byForward  bool
	hasNext    bool
	hasPrev    bool
}

// IsOutDto is a marker for out dto.
func (o ArticlesOutDto) IsOutDto() {}

// Articles returns articles.
func (o ArticlesOutDto) Articles() []ArticleTag {
	return o.articles
}

// HasNext returns true if next page exists otherwise false.
func (o ArticlesOutDto) HasNext() bool {
	return o.hasNext
}

// HasPrev returns true if prev page exists otherwise false.
func (o ArticlesOutDto) HasPrev() bool {
	return o.hasPrev
}

// ByForward returns true if this dto was fetched by next paging otherwise false.
func (o ArticlesOutDto) ByForward() bool {
	return o.byForward
}

// ByBackward returns true if this dto was fetched by prev paging otherwise false.
func (o ArticlesOutDto) ByBackward() bool {
	return o.byBackward
}

// ArticlesOutDtoOption is an option for ArticlesOutDto.
type ArticlesOutDtoOption func(*ArticlesOutDto)

// ArticlesOutDtoWithHasNext is an option for ArticlesOutDto.
func ArticlesOutDtoWithHasNext(hasNext bool) ArticlesOutDtoOption {
	return func(d *ArticlesOutDto) {
		if d.byBackward || d.hasPrev {
			return
		}
		d.hasNext = hasNext
		d.byForward = true
	}
}

// ArticlesOutDtoWithHasPrev is an option for ArticlesOutDto.
func ArticlesOutDtoWithHasPrev(hasPrev bool) ArticlesOutDtoOption {
	return func(d *ArticlesOutDto) {
		if d.byForward || d.hasNext {
			return
		}
		d.hasPrev = hasPrev
		d.byBackward = true
	}
}

// NewArticlesOutDto constructor of ArticlesOutDto.
func NewArticlesOutDto(articles []ArticleTag, options ...ArticlesOutDtoOption) ArticlesOutDto {
	d := ArticlesOutDto{
		articles: articles,
	}
	for _, option := range options {
		option(&d)
	}
	return d
}

// TagInDto is a dto for tag input.
type TagInDto struct {
	id string
}

// IsInDto is a marker for in dto.
func (i TagInDto) IsInDto() {}

// Id returns id.
func (i TagInDto) Id() string {
	return i.id
}

// NewTagInDto constructor of TagInDto.
func NewTagInDto(id string) TagInDto {
	return TagInDto{
		id: id,
	}
}

// TagsInDto is a dto for articles.
type TagsInDto struct {
	first  int
	last   int
	after  string
	before string
}

// IsInDto is a marker for in dto.
func (i TagsInDto) IsInDto() {}

func (i TagsInDto) First() int {
	return i.first
}

func (i TagsInDto) Last() int {
	return i.last
}

func (i TagsInDto) After() string {
	return i.after
}

func (i TagsInDto) Before() string {
	return i.before
}

type TagsInDtoOption func(*TagsInDto) error

// TagsInWithFirst specifies how many articles to retrieve from the beginning.
func TagsInWithFirst(first int) TagsInDtoOption {
	return func(d *TagsInDto) error {
		if d.last != 0 {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if last is set, first cannot be set.")
		}
		if d.before != "" {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if before is set, first cannot be set.")
		}
		d.first = first
		return nil
	}
}

// TagsInWithLast specifies how many articles to retrieve from the end.
func TagsInWithLast(last int) TagsInDtoOption {
	return func(d *TagsInDto) error {
		if d.first != 0 {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if first is set, last cannot be set.")
		}
		if d.after != "" {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if after is set, last cannot be set.")
		}
		d.last = last
		return nil
	}
}

// TagsInWithAfter specifies which cursor to get as the starting point.
//
// NOTE: must always be executed after TagsInWithFirst
// NOTE: if TagsInWithLast or TagsInWithBefore was executed, it will be returned ErrInvalidateTagsInDto.
func TagsInWithAfter(after string) TagsInDtoOption {
	return func(d *TagsInDto) error {
		if d.last != 0 {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if last is set, after cannot be set.")
		}
		if d.before != "" {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if before is set, after cannot be set.")
		}
		if d.first == 0 {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if first is not set, after cannot be set.")
		}
		d.after = after
		return nil
	}
}

// TagsInWithBefore specifies which cursor to get as the starting point.
//
// NOTE: must always be executed after TagsInWithLast
// NOTE: if TagsInWithFirst or TagsInWithAfter was executed, it will be returned ErrInvalidateTagsInDto.
func TagsInWithBefore(before string) TagsInDtoOption {
	return func(d *TagsInDto) error {
		if d.first != 0 {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if first is set, before cannot be set.")
		}
		if d.after != "" {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if after is set, before cannot be set.")
		}
		if d.last == 0 {
			return errors.WithMessage(ErrInvalidateTagsInDto, "if last is not set, before cannot be set.")
		}
		d.before = before
		return nil
	}
}

// NewTagsInDto constructor of TagsInDto.
// if options are invalid, return ErrInvalidateTagsInDto.
func NewTagsInDto(options ...TagsInDtoOption) (TagsInDto, error) {
	d := TagsInDto{}
	for _, option := range options {
		err := option(&d)
		if err != nil {
			return TagsInDto{}, err
		}
	}
	return d, nil
}

type TagArticle struct {
	Tag
	articles []Article
}

// IsOutDto is a marker for out dto.
func (t TagArticle) IsOutDto() {}

// Id returns id.
func (t TagArticle) Id() string {
	return t.id
}

// Name returns name.
func (t TagArticle) Name() string {
	return t.name
}

// Articles returns articles.
func (t TagArticle) Articles() []Article {
	return t.articles
}

// NewTagArticle constructor of TagArticle.
func NewTagArticle(id, name string, articles []Article) TagArticle {
	return TagArticle{
		Tag:      NewTag(id, name),
		articles: articles,
	}
}

// TagOutDto is a dto for tag output.
type TagOutDto struct {
	tag TagArticle
}

// IsOutDto is a marker for out dto.
func (o TagOutDto) IsOutDto() {}

// Tag returns tag.
func (o TagOutDto) Tag() TagArticle {
	return o.tag
}

// NewTagOutDto constructor of TagOutDto.
func NewTagOutDto(tag TagArticle) TagOutDto {
	return TagOutDto{
		tag: tag,
	}
}

// TagsOutDto is a dto for tags output.
type TagsOutDto struct {
	tags       []TagArticle
	byBackward bool
	byForward  bool
	hasNext    bool
	hasPrev    bool
}

// IsOutDto is a marker for out dto.
func (o TagsOutDto) IsOutDto() {}

// Tags returns tags.
func (o TagsOutDto) Tags() []TagArticle {
	return o.tags
}

// HasNext returns true if next page exists otherwise false.
func (o TagsOutDto) HasNext() bool {
	return o.hasNext
}

// HasPrev returns true if prev page exists otherwise false.
func (o TagsOutDto) HasPrev() bool {
	return o.hasPrev
}

// ByForward returns true if this dto was fetched by next paging otherwise false.
func (o TagsOutDto) ByForward() bool {
	return o.byForward
}

// ByBackward returns true if this dto was fetched by prev paging otherwise false.
func (o TagsOutDto) ByBackward() bool {
	return o.byBackward
}

// TagsOutDtoOption is an option for ArticlesOutDto.
type TagsOutDtoOption func(*TagsOutDto)

// TagsOutDtoWithHasNext is an option for ArticlesOutDto.
func TagsOutDtoWithHasNext(hasNext bool) TagsOutDtoOption {
	return func(o *TagsOutDto) {
		if o.byBackward || o.hasPrev {
			return
		}
		o.hasNext = hasNext
		o.byForward = true
	}
}

// TagsOutDtoWithHasPrev is an option for ArticlesOutDto.
func TagsOutDtoWithHasPrev(hasPrev bool) TagsOutDtoOption {
	return func(o *TagsOutDto) {
		if o.byForward || o.hasNext {
			return
		}
		o.hasPrev = hasPrev
		o.byBackward = true
	}
}

// NewTagsOutDto constructor of TagsOutDto.
func NewTagsOutDto(tags []TagArticle, options ...TagsOutDtoOption) TagsOutDto {
	o := TagsOutDto{
		tags: tags,
	}
	for _, option := range options {
		option(&o)
	}
	return o
}
