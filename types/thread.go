package types

type Thread struct {
	No            int    `json:"no"`
	Resto         int    `json:"resto"`
	Sticky        int    `json:"sticky"`
	Closed        int    `json:"closed"`
	Now           string `json:"now"`
	Time          int    `json:"time"`
	Name          string `json:"name"`
	Trip          string `json:"trip"`
	Id            string `json:"id"`
	Capcode       string `json:"capcode"`
	Country       string `json:"country"`
	CountryName   string `json:"country_name"`
	Sub           string `json:"sub"`
	Com           string `json:"com"`
	Tim           int    `json:"tim"`
	Filename      string `json:"filename"`
	Ext           string `json:"ext"`
	Fsize         int    `json:"fsize"`
	Md5           string `json:"md5"`
	W             int    `json:"w"`
	H             int    `json:"h"`
	TnW           int    `json:"tn_w"`
	TnH           int    `json:"tn_h"`
	Filedeleted   int    `json:"filedeleted"`
	Spoiler       int    `json:"spoiler"`
	CustomSpoiler int    `json:"custom_spoiler"`
	OmittedPosts  int    `json:"omitted_posts"`
	OmittedImages int    `json:"omitted_images"`
	Replies       int    `json:"replies"`
	Images        int    `json:"images"`
	Bumplimit     int    `json:"bumplimit"`
	Imagelimit    int    `json:"imagelimit"`
	Tag           string `json:"tag"`
	SematicUrl    string `json:"semantic_url"`
	Since4pass    int    `json:"since4pass"`
	UniqueIps     int    `json:"unique_ips"`
	MImg          int    `json:"m_img"`
	Archived      int    `json:"archived"`
	ArchivedOn    int    `json:"archived_on"`
}

type Page struct {
	Page    int      `json:"page"`
	Threads []Thread `json:"threads"`
}
