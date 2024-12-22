package dto

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"io"
	"net/url"
)

var (
	ErrInvalidateArticlesInDTO = errors.New("invalidate articles in dto")
	ErrInvalidateTagsInDTO     = errors.New("invalidate tags in dto")
)

type ArticleInDTO struct {
	id string
}

// ID returns id.
func (i ArticleInDTO) ID() string {
	return i.id
}

// IsInDTO is a marker for in dto.
func (i ArticleInDTO) IsInDTO() {}

// NewArticleInDTO constructor of ArticleInDTO.
func NewArticleInDTO(id string) ArticleInDTO {
	return ArticleInDTO{
		id: id,
	}
}

// ArticlesInDTO is a dto for articles.
type ArticlesInDTO struct {
	first  int
	last   int
	after  string
	before string
}

func (i ArticlesInDTO) First() int {
	return i.first
}

func (i ArticlesInDTO) Last() int {
	return i.last
}

func (i ArticlesInDTO) After() string {
	return i.after
}

func (i ArticlesInDTO) Before() string {
	return i.before
}

type ArticlesInDTOOption func(*ArticlesInDTO) error

// ArticlesInWithFirst specifies how many articles to retrieve from the beginning.
func ArticlesInWithFirst(first int) ArticlesInDTOOption {
	return func(d *ArticlesInDTO) error {
		if d.last != 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if last is set, first cannot be set.")
		}
		if d.before != "" {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if before is set, first cannot be set.")
		}
		d.first = first
		return nil
	}
}

// ArticlesInWithLast specifies how many articles to retrieve from the end.
func ArticlesInWithLast(last int) ArticlesInDTOOption {
	return func(d *ArticlesInDTO) error {
		if d.first != 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if first is set, last cannot be set.")
		}
		if d.after != "" {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if after is set, last cannot be set.")
		}
		d.last = last
		return nil
	}
}

// ArticlesInWithAfter specifies which cursor to get as the starting point.
//
// NOTE: must always be executed after ArticlesInWithFirst
// NOTE: if ArticlesInWithLast or ArticlesInWithBefore was executed, it will be returned ErrInvalidateArticlesInDTO.
func ArticlesInWithAfter(after string) ArticlesInDTOOption {
	return func(d *ArticlesInDTO) error {
		if d.last != 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if last is set, after cannot be set.")
		}
		if d.before != "" {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if before is set, after cannot be set.")
		}
		if d.first == 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if first is not set, after cannot be set.")
		}
		d.after = after
		return nil
	}
}

// ArticlesInWithBefore specifies which cursor to get as the starting point.
//
// NOTE: must always be executed after ArticlesInWithLast
// NOTE: if ArticlesInWithFirst or ArticlesInWithAfter was executed, it will be returned ErrInvalidateArticlesInDTO.
func ArticlesInWithBefore(before string) ArticlesInDTOOption {
	return func(d *ArticlesInDTO) error {
		if d.first != 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if first is set, before cannot be set.")
		}
		if d.after != "" {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if after is set, before cannot be set.")
		}
		if d.last == 0 {
			return errors.WithMessage(ErrInvalidateArticlesInDTO, "if last is not set, before cannot be set.")
		}
		d.before = before
		return nil
	}
}

// NewArticlesInDTO constructor of ArticlesInDTO.
// if options are invalid, return ErrInvalidateArticlesInDTO.
func NewArticlesInDTO(options ...ArticlesInDTOOption) (ArticlesInDTO, error) {
	d := ArticlesInDTO{}
	for _, option := range options {
		err := option(&d)
		if err != nil {
			return ArticlesInDTO{}, err
		}
	}
	return d, nil
}

type Article struct {
	id           string
	title        string
	body         string
	thumbnailURL url.URL
	createdAt    synchro.Time[tz.UTC]
	updatedAt    synchro.Time[tz.UTC]
}

// IsOutDTO is a marker for out dto.
func (a Article) IsOutDTO() {}

// ID returns id.
func (a Article) ID() string {
	return a.id
}

// Title returns title.
func (a Article) Title() string {
	return a.title
}

// ThumbnailURL returns thumbnail url.
func (a Article) ThumbnailURL() url.URL {
	return a.thumbnailURL
}

// CreatedAt returns created at.
func (a Article) CreatedAt() synchro.Time[tz.UTC] {
	return a.createdAt
}

// UpdatedAt returns updated at.
func (a Article) UpdatedAt() synchro.Time[tz.UTC] {
	return a.updatedAt
}

func NewArticle(id, title, body string, thumbnailURL url.URL, createdAt, updatedAt synchro.Time[tz.UTC]) Article {
	return Article{
		id:           id,
		title:        title,
		body:         body,
		thumbnailURL: thumbnailURL,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

type Tag struct {
	id   string
	name string
}

// IsOutDTO is a marker for out dto.
func (t Tag) IsOutDTO() {}

func (t Tag) ID() string {
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

func NewArticleTag(id, title, body string, thumbnailURL url.URL, createdAt, updatedAt synchro.Time[tz.UTC], tags []Tag) ArticleTag {
	return ArticleTag{
		Article: NewArticle(id, title, body, thumbnailURL, createdAt, updatedAt),
		tags:    tags,
	}
}

type ArticleOutDTO struct {
	article ArticleTag
}

func NewArticleOutDTO(article ArticleTag) ArticleOutDTO {
	return ArticleOutDTO{
		article: article,
	}
}

func (o ArticleOutDTO) Article() ArticleTag {
	return o.article
}

// IsOutDTO is a marker for out dto.
func (o ArticleOutDTO) IsOutDTO() {}

// ArticlesOutDTO is a dto for articles.
type ArticlesOutDTO struct {
	articles   []ArticleTag
	byBackward bool
	byForward  bool
	hasNext    bool
	hasPrev    bool
}

// IsOutDTO is a marker for out dto.
func (o ArticlesOutDTO) IsOutDTO() {}

// Articles returns articles.
func (o ArticlesOutDTO) Articles() []ArticleTag {
	return o.articles
}

// HasNext returns true if next page exists otherwise false.
func (o ArticlesOutDTO) HasNext() bool {
	return o.hasNext
}

// HasPrev returns true if prev page exists otherwise false.
func (o ArticlesOutDTO) HasPrev() bool {
	return o.hasPrev
}

// ByForward returns true if this dto was fetched by next paging otherwise false.
func (o ArticlesOutDTO) ByForward() bool {
	return o.byForward
}

// ByBackward returns true if this dto was fetched by prev paging otherwise false.
func (o ArticlesOutDTO) ByBackward() bool {
	return o.byBackward
}

// ArticlesOutDTOOption is an option for ArticlesOutDTO.
type ArticlesOutDTOOption func(*ArticlesOutDTO)

// ArticlesOutDTOWithHasNext is an option for ArticlesOutDTO.
func ArticlesOutDTOWithHasNext(hasNext bool) ArticlesOutDTOOption {
	return func(d *ArticlesOutDTO) {
		if d.byBackward || d.hasPrev {
			return
		}
		d.hasNext = hasNext
		d.byForward = true
	}
}

// ArticlesOutDTOWithHasPrev is an option for ArticlesOutDTO.
func ArticlesOutDTOWithHasPrev(hasPrev bool) ArticlesOutDTOOption {
	return func(d *ArticlesOutDTO) {
		if d.byForward || d.hasNext {
			return
		}
		d.hasPrev = hasPrev
		d.byBackward = true
	}
}

// NewArticlesOutDTO constructor of ArticlesOutDTO.
func NewArticlesOutDTO(articles []ArticleTag, options ...ArticlesOutDTOOption) ArticlesOutDTO {
	d := ArticlesOutDTO{
		articles: articles,
	}
	for _, option := range options {
		option(&d)
	}
	return d
}

// TagInDTO is a dto for tag input.
type TagInDTO struct {
	id string
}

// IsInDTO is a marker for in dto.
func (i TagInDTO) IsInDTO() {}

// ID returns id.
func (i TagInDTO) ID() string {
	return i.id
}

// NewTagInDTO constructor of TagInDTO.
func NewTagInDTO(id string) TagInDTO {
	return TagInDTO{
		id: id,
	}
}

// TagsInDTO is a dto for articles.
type TagsInDTO struct {
	first  int
	last   int
	after  string
	before string
}

// IsInDTO is a marker for in dto.
func (i TagsInDTO) IsInDTO() {}

func (i TagsInDTO) First() int {
	return i.first
}

func (i TagsInDTO) Last() int {
	return i.last
}

func (i TagsInDTO) After() string {
	return i.after
}

func (i TagsInDTO) Before() string {
	return i.before
}

type TagsInDTOOption func(*TagsInDTO) error

// TagsInWithFirst specifies how many articles to retrieve from the beginning.
func TagsInWithFirst(first int) TagsInDTOOption {
	return func(d *TagsInDTO) error {
		if d.last != 0 {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if last is set, first cannot be set.")
		}
		if d.before != "" {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if before is set, first cannot be set.")
		}
		d.first = first
		return nil
	}
}

// TagsInWithLast specifies how many articles to retrieve from the end.
func TagsInWithLast(last int) TagsInDTOOption {
	return func(d *TagsInDTO) error {
		if d.first != 0 {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if first is set, last cannot be set.")
		}
		if d.after != "" {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if after is set, last cannot be set.")
		}
		d.last = last
		return nil
	}
}

// TagsInWithAfter specifies which cursor to get as the starting point.
//
// NOTE: must always be executed after TagsInWithFirst
// NOTE: if TagsInWithLast or TagsInWithBefore was executed, it will be returned ErrInvalidateTagsInDTO.
func TagsInWithAfter(after string) TagsInDTOOption {
	return func(d *TagsInDTO) error {
		if d.last != 0 {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if last is set, after cannot be set.")
		}
		if d.before != "" {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if before is set, after cannot be set.")
		}
		if d.first == 0 {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if first is not set, after cannot be set.")
		}
		d.after = after
		return nil
	}
}

// TagsInWithBefore specifies which cursor to get as the starting point.
//
// NOTE: must always be executed after TagsInWithLast
// NOTE: if TagsInWithFirst or TagsInWithAfter was executed, it will be returned ErrInvalidateTagsInDTO.
func TagsInWithBefore(before string) TagsInDTOOption {
	return func(d *TagsInDTO) error {
		if d.first != 0 {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if first is set, before cannot be set.")
		}
		if d.after != "" {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if after is set, before cannot be set.")
		}
		if d.last == 0 {
			return errors.WithMessage(ErrInvalidateTagsInDTO, "if last is not set, before cannot be set.")
		}
		d.before = before
		return nil
	}
}

// NewTagsInDTO constructor of TagsInDTO.
// if options are invalid, return ErrInvalidateTagsInDTO.
func NewTagsInDTO(options ...TagsInDTOOption) (TagsInDTO, error) {
	d := TagsInDTO{}
	for _, option := range options {
		err := option(&d)
		if err != nil {
			return TagsInDTO{}, err
		}
	}
	return d, nil
}

type TagArticle struct {
	Tag
	articles []Article
}

// IsOutDTO is a marker for out dto.
func (t TagArticle) IsOutDTO() {}

// ID returns id.
func (t TagArticle) ID() string {
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

// TagOutDTO is a dto for tag output.
type TagOutDTO struct {
	tag TagArticle
}

// IsOutDTO is a marker for out dto.
func (o TagOutDTO) IsOutDTO() {}

// Tag returns tag.
func (o TagOutDTO) Tag() TagArticle {
	return o.tag
}

// NewTagOutDTO constructor of TagOutDTO.
func NewTagOutDTO(tag TagArticle) TagOutDTO {
	return TagOutDTO{
		tag: tag,
	}
}

// TagsOutDTO is a dto for tags output.
type TagsOutDTO struct {
	tags       []TagArticle
	byBackward bool
	byForward  bool
	hasNext    bool
	hasPrev    bool
}

// IsOutDTO is a marker for out dto.
func (o TagsOutDTO) IsOutDTO() {}

// Tags returns tags.
func (o TagsOutDTO) Tags() []TagArticle {
	return o.tags
}

// HasNext returns true if next page exists otherwise false.
func (o TagsOutDTO) HasNext() bool {
	return o.hasNext
}

// HasPrev returns true if prev page exists otherwise false.
func (o TagsOutDTO) HasPrev() bool {
	return o.hasPrev
}

// ByForward returns true if this dto was fetched by next paging otherwise false.
func (o TagsOutDTO) ByForward() bool {
	return o.byForward
}

// ByBackward returns true if this dto was fetched by prev paging otherwise false.
func (o TagsOutDTO) ByBackward() bool {
	return o.byBackward
}

// TagsOutDTOOption is an option for ArticlesOutDTO.
type TagsOutDTOOption func(*TagsOutDTO)

// TagsOutDTOWithHasNext is an option for ArticlesOutDTO.
func TagsOutDTOWithHasNext(hasNext bool) TagsOutDTOOption {
	return func(o *TagsOutDTO) {
		if o.byBackward || o.hasPrev {
			return
		}
		o.hasNext = hasNext
		o.byForward = true
	}
}

// TagsOutDTOWithHasPrev is an option for ArticlesOutDTO.
func TagsOutDTOWithHasPrev(hasPrev bool) TagsOutDTOOption {
	return func(o *TagsOutDTO) {
		if o.byForward || o.hasNext {
			return
		}
		o.hasPrev = hasPrev
		o.byBackward = true
	}
}

// NewTagsOutDTO constructor of TagsOutDTO.
func NewTagsOutDTO(tags []TagArticle, options ...TagsOutDTOOption) TagsOutDTO {
	o := TagsOutDTO{
		tags: tags,
	}
	for _, option := range options {
		option(&o)
	}
	return o
}

// CreateArticleInDTO is a dto for creating an article.
type CreateArticleInDTO struct {
	title            string
	body             string
	thumbnailURL     url.URL
	tagNames         []string
	clientMutationID string
}

// IsInDTO is a marker for in dto.
func (a CreateArticleInDTO) IsInDTO() {}

// Title returns title.
func (a CreateArticleInDTO) Title() string {
	return a.title
}

// Body returns body.
func (a CreateArticleInDTO) Body() string {
	return a.body
}

// ThumbnailURL returns thumbnail url.
func (a CreateArticleInDTO) ThumbnailURL() url.URL {
	return a.thumbnailURL
}

// TagNames returns tag names.
func (a CreateArticleInDTO) TagNames() []string {
	return a.tagNames
}

// ClientMutationID returns client mutation id.
func (a CreateArticleInDTO) ClientMutationID() string {
	return a.clientMutationID
}

// NewCreateArticleInDTO constructor of CreateArticleInDTO.
func NewCreateArticleInDTO(title, body string, thumbnailURL url.URL, tagNames []string, clientMutationID string) CreateArticleInDTO {
	return CreateArticleInDTO{
		title:            title,
		body:             body,
		thumbnailURL:     thumbnailURL,
		tagNames:         tagNames,
		clientMutationID: clientMutationID,
	}
}

// CreateArticleOutDTO is a dto for creating an article.
type CreateArticleOutDTO struct {
	eventID          string
	articleID        string
	clientMutationID string
}

// IsOutDTO is a marker for out dto.
func (a CreateArticleOutDTO) IsOutDTO() {}

// EventID returns event id.
func (a CreateArticleOutDTO) EventID() string {
	return a.eventID
}

// ArticleID returns article id.
func (a CreateArticleOutDTO) ArticleID() string {
	return a.articleID
}

// ClientMutationID returns client mutation id.
func (a CreateArticleOutDTO) ClientMutationID() string {
	return a.clientMutationID
}

// NewCreateArticleOutDTO constructor of CreateArticleOutDTO.
func NewCreateArticleOutDTO(eventID, articleID, clientMutationID string) CreateArticleOutDTO {
	return CreateArticleOutDTO{
		eventID:          eventID,
		articleID:        articleID,
		clientMutationID: clientMutationID,
	}
}

// UpdateArticleTitleInDTO is a dto for updating an article.
type UpdateArticleTitleInDTO struct {
	id               string
	title            string
	clientMutationID string
}

// IsInDTO is a marker for in dto.
func (u UpdateArticleTitleInDTO) IsInDTO() {}

// ID returns id.
func (u UpdateArticleTitleInDTO) ID() string {
	return u.id
}

// Title returns title.
func (u UpdateArticleTitleInDTO) Title() string {
	return u.title
}

// ClientMutationID returns client mutation id.
func (u UpdateArticleTitleInDTO) ClientMutationID() string {
	return u.clientMutationID
}

// NewUpdateArticleTitleInDTO constructor of UpdateArticleTitleInDTO.
func NewUpdateArticleTitleInDTO(id, title, clientMutationID string) UpdateArticleTitleInDTO {
	return UpdateArticleTitleInDTO{
		id:               id,
		title:            title,
		clientMutationID: clientMutationID,
	}
}

// UpdateArticleTitleOutDTO is a dto for updating an article.
type UpdateArticleTitleOutDTO struct {
	eventID          string
	articleID        string
	clientMutationID string
}

// IsOutDTO is a marker for out dto.
func (u UpdateArticleTitleOutDTO) IsOutDTO() {}

// EventID returns event id.
func (u UpdateArticleTitleOutDTO) EventID() string {
	return u.eventID
}

// ArticleID returns article id.
func (u UpdateArticleTitleOutDTO) ArticleID() string {
	return u.articleID
}

// ClientMutationID returns client mutation id.
func (u UpdateArticleTitleOutDTO) ClientMutationID() string {
	return u.clientMutationID
}

// NewUpdateArticleTitleOutDTO constructor of UpdateArticleTitleOutDTO.
func NewUpdateArticleTitleOutDTO(eventID, articleID, clientMutationID string) UpdateArticleTitleOutDTO {
	return UpdateArticleTitleOutDTO{
		eventID:          eventID,
		articleID:        articleID,
		clientMutationID: clientMutationID,
	}
}

// UpdateArticleBodyInDTO is a dto for updating an article content.
type UpdateArticleBodyInDTO struct {
	id               string
	content          string
	clientMutationID string
}

// IsInDTO is a marker for in dto.
func (u UpdateArticleBodyInDTO) IsInDTO() {}

// ID returns id.
func (u UpdateArticleBodyInDTO) ID() string {
	return u.id
}

// Content returns content.
func (u UpdateArticleBodyInDTO) Content() string {
	return u.content
}

// ClientMutationID returns client mutation id.
func (u UpdateArticleBodyInDTO) ClientMutationID() string {
	return u.clientMutationID
}

// NewUpdateArticleBodyInDTO constructor of UpdateArticleBodyInDTO.
func NewUpdateArticleBodyInDTO(id, content, clientMutationID string) UpdateArticleBodyInDTO {
	return UpdateArticleBodyInDTO{
		id:               id,
		content:          content,
		clientMutationID: clientMutationID,
	}
}

// UpdateArticleBodyOutDTO is a dto for updating an article content.
type UpdateArticleBodyOutDTO struct {
	eventID          string
	articleID        string
	clientMutationID string
}

// IsOutDTO is a marker for out dto.
func (u UpdateArticleBodyOutDTO) IsOutDTO() {}

// EventID returns event id.
func (u UpdateArticleBodyOutDTO) EventID() string {
	return u.eventID
}

// ArticleID returns article id.
func (u UpdateArticleBodyOutDTO) ArticleID() string {
	return u.articleID
}

// ClientMutationID returns client mutation id.
func (u UpdateArticleBodyOutDTO) ClientMutationID() string {
	return u.clientMutationID
}

// NewUpdateArticleBodyOutDTO constructor of UpdateArticleBodyOutDTO.
func NewUpdateArticleBodyOutDTO(eventID, articleID, clientMutationID string) UpdateArticleBodyOutDTO {
	return UpdateArticleBodyOutDTO{
		eventID:          eventID,
		articleID:        articleID,
		clientMutationID: clientMutationID,
	}
}

// UpdateArticleThumbnailInDTO is a dto for updating an article thumbnail.
type UpdateArticleThumbnailInDTO struct {
	id               string
	thumbnail        url.URL
	clientMutationID string
}

// ID returns id.
func (u UpdateArticleThumbnailInDTO) ID() string {
	return u.id
}

// Thumbnail returns thumbnail.
func (u UpdateArticleThumbnailInDTO) Thumbnail() url.URL {
	return u.thumbnail
}

// ClientMutationID returns client mutation id.
func (u UpdateArticleThumbnailInDTO) ClientMutationID() string {
	return u.clientMutationID
}

// NewUpdateArticleThumbnailInDTO constructor of UpdateArticleThumbnailInDTO.
func NewUpdateArticleThumbnailInDTO(id string, thumbnail url.URL, clientMutationID string) UpdateArticleThumbnailInDTO {
	return UpdateArticleThumbnailInDTO{
		id:               id,
		thumbnail:        thumbnail,
		clientMutationID: clientMutationID,
	}
}

// UpdateArticleThumbnailOutDTO is a dto for updating an article thumbnail.
type UpdateArticleThumbnailOutDTO struct {
	eventID          string
	articleID        string
	clientMutationID string
}

// EventID returns event id.
func (u UpdateArticleThumbnailOutDTO) EventID() string {
	return u.eventID
}

// ArticleID returns article id.
func (u UpdateArticleThumbnailOutDTO) ArticleID() string {
	return u.articleID
}

// ClientMutationID returns client mutation id.
func (u UpdateArticleThumbnailOutDTO) ClientMutationID() string {
	return u.clientMutationID
}

// NewUpdateArticleThumbnailOutDTO constructor of UpdateArticleThumbnailOutDTO.
func NewUpdateArticleThumbnailOutDTO(eventID, articleID, clientMutationID string) UpdateArticleThumbnailOutDTO {
	return UpdateArticleThumbnailOutDTO{
		eventID:          eventID,
		articleID:        articleID,
		clientMutationID: clientMutationID,
	}
}

type AttachTagsInDTO struct {
	id               string
	tagNames         []string
	clientMutationID string
}

// ID returns id.
func (a AttachTagsInDTO) ID() string {
	return a.id
}

// TagNames returns tag names.
func (a AttachTagsInDTO) TagNames() []string {
	return a.tagNames
}

// ClientMutationID returns client mutation id.
func (a AttachTagsInDTO) ClientMutationID() string {
	return a.clientMutationID
}

// NewAttachTagsInDTO constructor of AttachTagsInDTO.
func NewAttachTagsInDTO(id string, tagNames []string, clientMutationID string) AttachTagsInDTO {
	return AttachTagsInDTO{
		id:               id,
		tagNames:         tagNames,
		clientMutationID: clientMutationID,
	}
}

// AttachTagsOutDTO is a dto for attaching tags to an article.
type AttachTagsOutDTO struct {
	eventID          string
	articleID        string
	clientMutationID string
}

// EventID returns event id.
func (a AttachTagsOutDTO) EventID() string {
	return a.eventID
}

// ArticleID returns article id.
func (a AttachTagsOutDTO) ArticleID() string {
	return a.articleID
}

// ClientMutationID returns client mutation id.
func (a AttachTagsOutDTO) ClientMutationID() string {
	return a.clientMutationID
}

// NewAttachTagsOutDTO constructor of AttachTagsOutDTO.
func NewAttachTagsOutDTO(eventID, articleID, clientMutationID string) AttachTagsOutDTO {
	return AttachTagsOutDTO{
		eventID:          eventID,
		articleID:        articleID,
		clientMutationID: clientMutationID,
	}
}

type DetachTagsInDTO struct {
	id               string
	tagNames         []string
	clientMutationID string
}

// ID returns id.
func (a DetachTagsInDTO) ID() string {
	return a.id
}

// TagNames returns tag names.
func (a DetachTagsInDTO) TagNames() []string {
	return a.tagNames
}

// ClientMutationID returns client mutation id.
func (a DetachTagsInDTO) ClientMutationID() string {
	return a.clientMutationID
}

// NewDetachTagsInDTO constructor of DetachTagsInDTO.
func NewDetachTagsInDTO(id string, tagNames []string, clientMutationID string) DetachTagsInDTO {
	return DetachTagsInDTO{
		id:               id,
		tagNames:         tagNames,
		clientMutationID: clientMutationID,
	}
}

// DetachTagsOutDTO is a dto for attaching tags to an article.
type DetachTagsOutDTO struct {
	eventID          string
	articleID        string
	clientMutationID string
}

// EventID returns event id.
func (a DetachTagsOutDTO) EventID() string {
	return a.eventID
}

// ArticleID returns article id.
func (a DetachTagsOutDTO) ArticleID() string {
	return a.articleID
}

// ClientMutationID returns client mutation id.
func (a DetachTagsOutDTO) ClientMutationID() string {
	return a.clientMutationID
}

// NewDetachTagsOutDTO constructor of DetachTagsOutDTO.
func NewDetachTagsOutDTO(eventID, articleID, clientMutationID string) DetachTagsOutDTO {
	return DetachTagsOutDTO{
		eventID:          eventID,
		articleID:        articleID,
		clientMutationID: clientMutationID,
	}
}

// UploadImageInDTO is a dto for uploading an image.
type UploadImageInDTO struct {
	data             io.ReadSeeker
	filename         string
	clientMutationID string
}

// Data returns data.
func (u UploadImageInDTO) Data() io.ReadSeeker {
	return u.data
}

// Filename returns filename.
func (u UploadImageInDTO) Filename() string {
	return u.filename
}

// ClientMutationID returns client mutation id.
func (u UploadImageInDTO) ClientMutationID() string {
	return u.clientMutationID
}

// NewUploadImageInDTO constructor of UploadImageInDTO.
func NewUploadImageInDTO(data io.ReadSeeker, filename, clientMutationID string) UploadImageInDTO {
	return UploadImageInDTO{
		data:             data,
		filename:         filename,
		clientMutationID: clientMutationID,
	}
}

// UploadImageOutDTO is a dto for uploading an image.
type UploadImageOutDTO struct {
	imageURL         url.URL
	clientMutationID string
}

// ImageURL returns image url.
func (u UploadImageOutDTO) ImageURL() url.URL {
	return u.imageURL
}

// ClientMutationID returns client mutation id.
func (u UploadImageOutDTO) ClientMutationID() string {
	return u.clientMutationID
}

// NewUploadImageOutDTO constructor of UploadImageOutDTO.
func NewUploadImageOutDTO(imageURL url.URL, clientMutationID string) UploadImageOutDTO {
	return UploadImageOutDTO{
		imageURL:         imageURL,
		clientMutationID: clientMutationID,
	}
}
