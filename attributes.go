package html_types

import "time"

// Global attributes
type GlobalAttributes struct {
    AccessKey         string             `json:"accesskey,omitempty"`
    AutoCapitalize    string             `json:"autocapitalize,omitempty"`
    Class             string             `json:"class,omitempty"`
    ContentEditable   bool               `json:"contenteditable,omitempty"`
    Data              map[string]string  `json:"data,omitempty"` // Data attributes in the form of data-*
    Dir               string             `json:"dir,omitempty"`
    Draggable         bool               `json:"draggable,omitempty"`
    EnterKeyHint      string             `json:"enterkeyhint,omitempty"`
    Hidden            bool               `json:"hidden,omitempty"`
    Id                string             `json:"id,omitempty"`
    InputMode         string             `json:"inputmode,omitempty"`
    Is                string             `json:"is,omitempty"`
    ItemId            string             `json:"itemid,omitempty"`
    ItemProp          string             `json:"itemprop,omitempty"`
    ItemRef           string             `json:"itemref,omitempty"`
    ItemScope         bool               `json:"itemscope,omitempty"`
    ItemType          string             `json:"itemtype,omitempty"`
    Lang              string             `json:"lang,omitempty"`
    Part              string             `json:"part,omitempty"`
    Slot              string             `json:"slot,omitempty"`
    SpellCheck        bool               `json:"spellcheck,omitempty"`
    Style             string             `json:"style,omitempty"`
    TabIndex          int                `json:"tabindex,omitempty"`
    Title             string             `json:"title,omitempty"`
    Translate         string             `json:"translate,omitempty"`
    Role              string             `json:"role,omitempty"`

    // Event handler attributes
    OnClick           string             `json:"onclick,omitempty"`
    OnChange          string             `json:"onchange,omitempty"`
    OnInput           string             `json:"oninput,omitempty"`
    OnSubmit          string             `json:"onsubmit,omitempty"`
    OnKeyDown         string             `json:"onkeydown,omitempty"`
    OnKeyPress        string             `json:"onkeypress,omitempty"`
    OnKeyUp           string             `json:"onkeyup,omitempty"`
    OnMouseDown       string             `json:"onmousedown,omitempty"`
    OnMouseMove       string             `json:"onmousemove,omitempty"`
    OnMouseOut        string             `json:"onmouseout,omitempty"`
    OnMouseOver       string             `json:"onmouseover,omitempty"`
    OnMouseUp         string             `json:"onmouseup,omitempty"`
}

// A attributes
type AnchorAttributes struct {
    GlobalAttributes
    Href              string             `json:"href,omitempty"`
    Target            string             `json:"target,omitempty"`
    Download          string             `json:"download,omitempty"`
    Ping              string             `json:"ping,omitempty"`
    Rel               string             `json:"rel,omitempty"`
    HrefLang          string             `json:"hreflang,omitempty"`
    ReferrerPolicy    string             `json:"referrerpolicy,omitempty"`
    Type              string             `json:"type,omitempty"`
}

// Area attributes
type AreaAttributes struct {
    GlobalAttributes
    Alt               string             `json:"alt,omitempty"`
    Coords            string             `json:"coords,omitempty"`
    Shape             string             `json:"shape,omitempty"`
    Href              string             `json:"href,omitempty"`
    Target            string             `json:"target,omitempty"`
    Download          string             `json:"download,omitempty"`
    Ping              string             `json:"ping,omitempty"`
    Rel               string             `json:"rel,omitempty"`
    HrefLang          string             `json:"hreflang,omitempty"`
    ReferrerPolicy    string             `json:"referrerpolicy,omitempty"`
}

// Audio, Video attributes
type MediaAttributes struct {
    GlobalAttributes
    Src               string             `json:"src,omitempty"`
    CrossOrigin       string             `json:"crossorigin,omitempty"`
    Preload           string             `json:"preload,omitempty"`
    AutoPlay          bool               `json:"autoplay,omitempty"`
    Controls          bool               `json:"controls,omitempty"`
    Loop              bool               `json:"loop,omitempty"`
    Muted             bool               `json:"muted,omitempty"`
    Default           bool               `json:"default,omitempty"`
    Kind              string             `json:"kind,omitempty"`
    Label             string             `json:"label,omitempty"`
    SrcLang           string             `json:"srclang,omitempty"`
}

// Base attributes
type BaseAttributes struct {
    GlobalAttributes
    Href              string             `json:"href,omitempty"`
    Target            string             `json:"target,omitempty"`
}

// Blockquote, Q attributes
type QuoteAttributes struct {
    GlobalAttributes
    Cite              string             `json:"cite,omitempty"`
}

// Button attributes
type ButtonAttributes struct {
    GlobalAttributes
    AutoFocus         bool               `json:"autofocus,omitempty"`
    Disabled          bool               `json:"disabled,omitempty"`
    Form              string             `json:"form,omitempty"`
    FormAction        string             `json:"formaction,omitempty"`
    FormEncType       string             `json:"formenctype,omitempty"`
    FormMethod        string             `json:"formmethod,omitempty"`
    FormNoValidate    bool               `json:"formnovalidate,omitempty"`
    FormTarget        string             `json:"formtarget,omitempty"`
    Name              string             `json:"name,omitempty"`
    Type              string             `json:"type,omitempty"`
    Value             string             `json:"value,omitempty"`
}

// Canvas attributes
type CanvasAttributes struct {
    GlobalAttributes
    Height            int                `json:"height,omitempty"`
    Width             int                `json:"width,omitempty"`
}

// Col, Colgroup attributes
type ColAttributes struct {
    GlobalAttributes
    Span              int                `json:"span,omitempty"`
}

// Data attributes
type DataAttributes struct {
    GlobalAttributes
    Value             string             `json:"value,omitempty"`
}

// Embed attributes
type EmbedAttributes struct {
    GlobalAttributes
    Src               string             `json:"src,omitempty"`
    Type              string             `json:"type,omitempty"`
    Width             int                `json:"width,omitempty"`
    Height            int                `json:"height,omitempty"`
}

// Fieldset attributes
type FieldsetAttributes struct {
    GlobalAttributes
    Disabled          bool               `json:"disabled,omitempty"`
    Form              string             `json:"form,omitempty"`
    Name              string             `json:"name,omitempty"`
}

// Form attributes
type FormAttributes struct {
    GlobalAttributes
    AcceptCharset     string             `json:"accept-charset,omitempty"`
    Action            string             `json:"action,omitempty"`
    AutoComplete      string             `json:"autocomplete,omitempty"`
    EncType           string             `json:"enctype,omitempty"`
    Method            string             `json:"method,omitempty"`
    NoValidate        bool               `json:"novalidate,omitempty"`
    Target            string             `json:"target,omitempty"`
}

// Iframe attributes
type IframeAttributes struct {
    GlobalAttributes
    Allow             string             `json:"allow,omitempty"`
    AllowFullScreen   bool               `json:"allowfullscreen,omitempty"`
    AllowPaymentRequest bool             `json:"allowpaymentrequest,omitempty"`
    FrameBorder       string             `json:"frameborder,omitempty"`
    Height            string             `json:"height,omitempty"`
    LongDesc          string             `json:"longdesc,omitempty"`
    MarginHeight      string             `json:"marginheight,omitempty"`
    MarginWidth       string             `json:"marginwidth,omitempty"`
    Name              string             `json:"name,omitempty"`
    ReferrerPolicy    string             `json:"referrerpolicy,omitempty"`
    Sandbox           string             `json:"sandbox,omitempty"`
    Scrolling         string             `json:"scrolling,omitempty"`
    Src               string             `json:"src,omitempty"`
    SrcDoc            string             `json:"srcdoc,omitempty"`
    Width             string             `json:"width,omitempty"`
}

// Img attributes
type ImgAttributes struct {
    GlobalAttributes
    Alt               string             `json:"alt,omitempty"`
    CrossOrigin       string             `json:"crossorigin,omitempty"`
    Decoding          string             `json:"decoding,omitempty"`
    Height            int                `json:"height,omitempty"`
    IsMap             bool               `json:"ismap,omitempty"`
    Loading           string             `json:"loading,omitempty"`
    ReferrerPolicy    string             `json:"referrerpolicy,omitempty"`
    Sizes             string             `json:"sizes,omitempty"`
    Src               string             `json:"src,omitempty"`
    SrcSet            string             `json:"srcset,omitempty"`
    UseMap            string             `json:"usemap,omitempty"`
    Width             int                `json:"width,omitempty"`
}

// Input attributes
type InputAttributes struct {
    GlobalAttributes
    Accept            string             `json:"accept,omitempty"`
    Alt               string             `json:"alt,omitempty"`
    AutoComplete      string             `json:"autocomplete,omitempty"`
    AutoFocus         bool               `json:"autofocus,omitempty"`
    Capture           string             `json:"capture,omitempty"`
    Checked           bool               `json:"checked,omitempty"`
    DirName           string             `json:"dirname,omitempty"`
    Disabled          bool               `json:"disabled,omitempty"`
    Form              string             `json:"form,omitempty"`
    FormAction        string             `json:"formaction,omitempty"`
    FormEncType       string             `json:"formenctype,omitempty"`
    FormMethod        string             `json:"formmethod,omitempty"`
    FormNoValidate    bool               `json:"formnovalidate,omitempty"`
    FormTarget        string             `json:"formtarget,omitempty"`
    Height            int                `json:"height,omitempty"`
    List              string             `json:"list,omitempty"`
    Max               string             `json:"max,omitempty"`
    MaxLength         int                `json:"maxlength,omitempty"`
    Min               string             `json:"min,omitempty"`
    MinLength         int                `json:"minlength,omitempty"`
    Multiple          bool               `json:"multiple,omitempty"`
    Name              string             `json:"name,omitempty"`
    Pattern           string             `json:"pattern,omitempty"`
    Placeholder       string             `json:"placeholder,omitempty"`
    ReadOnly          bool               `json:"readonly,omitempty"`
    Required          bool               `json:"required,omitempty"`
    Size              int                `json:"size,omitempty"`
    Src               string             `json:"src,omitempty"`
    Step              string             `json:"step,omitempty"`
    Type              string             `json:"type,omitempty"`
    Value             string             `json:"value,omitempty"`
    Width             int                `json:"width,omitempty"`
}

// Label attributes
type LabelAttributes struct {
    GlobalAttributes
    Form              string             `json:"form,omitempty"`
    HtmlFor           string             `json:"for,omitempty"`
}

// Link attributes
type LinkAttributes struct {
    GlobalAttributes
    CrossOrigin       string             `json:"crossorigin,omitempty"`
    Href              string             `json:"href,omitempty"`
    HrefLang          string             `json:"hreflang,omitempty"`
    Media             string             `json:"media,omitempty"`
    Rel               string             `json:"rel,omitempty"`
    Sizes             string             `json:"sizes,omitempty"`
    Type              string             `json:"type,omitempty"`
}

// Map attributes
type MapAttributes struct {
    GlobalAttributes
    Name              string             `json:"name,omitempty"`
}

// Meta attributes
type MetaAttributes struct {
    GlobalAttributes
    Charset           string             `json:"charset,omitempty"`
    Content           string             `json:"content,omitempty"`
    HttpEquiv         string             `json:"http-equiv,omitempty"`
    Name              string             `json:"name,omitempty"`
}

// Meter attributes
type MeterAttributes struct {
    GlobalAttributes
    Value             float64            `json:"value,omitempty"`
    Min               float64            `json:"min,omitempty"`
    Max               float64            `json:"max,omitempty"`
    Low               float64            `json:"low,omitempty"`
    High              float64            `json:"high,omitempty"`
    Optimum           float64            `json:"optimum,omitempty"`
}

// Object attributes
type ObjectAttributes struct {
    GlobalAttributes
    Data              string             `json:"data,omitempty"`
    Form              string             `json:"form,omitempty"`
    Height            int                `json:"height,omitempty"`
    Name              string             `json:"name,omitempty"`
    Type              string             `json:"type,omitempty"`
    UseMap            string             `json:"usemap,omitempty"`
    Width             int                `json:"width,omitempty"`
}

// Ol attributes
type OlAttributes struct {
    GlobalAttributes
    Reversed          bool               `json:"reversed,omitempty"`
    Start             int                `json:"start,omitempty"`
    Type              string             `json:"type,omitempty"`
}

// Optgroup attributes
type OptgroupAttributes struct {
    GlobalAttributes
    Disabled          bool               `json:"disabled,omitempty"`
    Label             string             `json:"label,omitempty"`
}

// Option attributes
type OptionAttributes struct {
    GlobalAttributes
    Disabled          bool               `json:"disabled,omitempty"`
    Label             string             `json:"label,omitempty"`
    Selected          bool               `json:"selected,omitempty"`
    Value             string             `json:"value,omitempty"`
}

// Output attributes
type OutputAttributes struct {
    GlobalAttributes
    For               string             `json:"for,omitempty"`
    Form              string             `json:"form,omitempty"`
    Name              string             `json:"name,omitempty"`
}

// Param attributes
type ParamAttributes struct {
    GlobalAttributes
    Name              string             `json:"name,omitempty"`
    Value             string             `json:"value,omitempty"`
}

// Progress attributes
type ProgressAttributes struct {
    GlobalAttributes
    Max               float64            `json:"max,omitempty"`
    Value             float64            `json:"value,omitempty"`
}

// Script attributes
type ScriptAttributes struct {
    GlobalAttributes
    Async             bool               `json:"async,omitempty"`
    Defer             bool               `json:"defer,omitempty"`
    Src               string             `json:"src,omitempty"`
    Type              string             `json:"type,omitempty"`
}

// Select attributes
type SelectAttributes struct {
    GlobalAttributes
    AutoComplete      string             `json:"autocomplete,omitempty"`
    AutoFocus         bool               `json:"autofocus,omitempty"`
    Disabled          bool               `json:"disabled,omitempty"`
    Form              string             `json:"form,omitempty"`
    Multiple          bool               `json:"multiple,omitempty"`
    Name              string             `json:"name,omitempty"`
    Required          bool               `json:"required,omitempty"`
    Size              int                `json:"size,omitempty"`
}

// Source attributes
type SourceAttributes struct {
    GlobalAttributes
    Media             string             `json:"media,omitempty"`
    Sizes             string             `json:"sizes,omitempty"`
    Src               string             `json:"src,omitempty"`
    SrcSet            string             `json:"srcset,omitempty"`
    Type              string             `json:"type,omitempty"`
}

// Style attributes
type StyleAttributes struct {
    GlobalAttributes
    Media             string             `json:"media,omitempty"`
    Type              string             `json:"type,omitempty"`
}

// Table attributes
type TableAttributes struct {
    GlobalAttributes
    Border            string             `json:"border,omitempty"`
}

// Textarea attributes
type TextareaAttributes struct {
    GlobalAttributes
    AutoComplete      string             `json:"autocomplete,omitempty"`
    AutoFocus         bool               `json:"autofocus,omitempty"`
    Cols              int                `json:"cols,omitempty"`
    DirName           string             `json:"dirname,omitempty"`
    Disabled          bool               `json:"disabled,omitempty"`
    Form              string             `json:"form,omitempty"`
    MaxLength         int                `json:"maxlength,omitempty"`
    MinLength         int                `json:"minlength,omitempty"`
    Name              string             `json:"name,omitempty"`
    Placeholder       string             `json:"placeholder,omitempty"`
    ReadOnly          bool               `json:"readonly,omitempty"`
    Required          bool               `json:"required,omitempty"`
    Rows              int                `json:"rows,omitempty"`
    Wrap              string             `json:"wrap,omitempty"`
}

// Time attributes
type TimeAttributes struct {
    GlobalAttributes
    DateTime          time.Time          `json:"datetime,omitempty"`
}

// Track attributes
type TrackAttributes struct {
    GlobalAttributes
    Default           bool               `json:"default,omitempty"`
    Kind              string             `json:"kind,omitempty"`
    Label             string             `json:"label,omitempty"`
    Src               string             `json:"src,omitempty"`
    SrcLang           string             `json:"srclang,omitempty"`
}

// Video attributes
type VideoAttributes struct {
    GlobalAttributes
    Src               string             `json:"src,omitempty"`
    CrossOrigin       string             `json:"crossorigin,omitempty"`
    Preload           string             `json:"preload,omitempty"`
    AutoPlay          bool               `json:"autoplay,omitempty"`
    Controls          bool               `json:"controls,omitempty"`
    Loop              bool               `json:"loop,omitempty"`
    Muted             bool               `json:"muted,omitempty"`
    Height            int                `json:"height,omitempty"`
    Width             int                `json:"width,omitempty"`
    Poster            string             `json:"poster,omitempty"`
}

// List attributes
type UlAttributes struct {
    GlobalAttributes
}

// List attributes
type LiAttributes struct {
    GlobalAttributes
    Value             int                `json:"value,omitempty"`
}

// Details attributes
type DetailsAttributes struct {
    GlobalAttributes
    Open              bool               `json:"open,omitempty"`
}

// Dialog attributes
type DialogAttributes struct {
    GlobalAttributes
    Open              bool               `json:"open,omitempty"`
}

// Abbr attributes
type AbbrAttributes struct {
    GlobalAttributes
}

// Address attributes
type AddressAttributes struct {
    GlobalAttributes
}

// Article attributes
type ArticleAttributes struct {
    GlobalAttributes
}

// Aside attributes
type AsideAttributes struct {
    GlobalAttributes
}

// B attributes
type BAttributes struct {
    GlobalAttributes
}

// Bdi attributes
type BdiAttributes struct {
    GlobalAttributes
}

// Bdo attributes
type BdoAttributes struct {
    GlobalAttributes
}

// Body attributes
type BodyAttributes struct {
    GlobalAttributes
    OnAfterPrint      string             `json:"onafterprint,omitempty"`
    OnBeforePrint     string             `json:"onbeforeprint,omitempty"`
    OnBeforeUnload    string             `json:"onbeforeunload,omitempty"`
    OnHashChange      string             `json:"onhashchange,omitempty"`
    OnLoad            string             `json:"onload,omitempty"`
    OnMessage         string             `json:"onmessage,omitempty"`
    OnOffline         string             `json:"onoffline,omitempty"`
    OnOnline          string             `json:"ononline,omitempty"`
    OnPageHide        string             `json:"onpagehide,omitempty"`
    OnPageShow        string             `json:"onpageshow,omitempty"`
    OnPopState        string             `json:"onpopstate,omitempty"`
    OnResize          string             `json:"onresize,omitempty"`
    OnStorage         string             `json:"onstorage,omitempty"`
    OnUnload          string             `json:"onunload,omitempty"`
}

// Br attributes
type BrAttributes struct {
    GlobalAttributes
}

// Caption attributes
type CaptionAttributes struct {
    GlobalAttributes
}

// Cite attributes
type CiteAttributes struct {
    GlobalAttributes
}

// Code attributes
type CodeAttributes struct {
    GlobalAttributes
}

// Datalist attributes
type DatalistAttributes struct {
    GlobalAttributes
}

// Dd attributes
type DdAttributes struct {
    GlobalAttributes
}

// Del attributes
type DelAttributes struct {
    GlobalAttributes
    Cite              string             `json:"cite,omitempty"`
    DateTime          string             `json:"datetime,omitempty"`
}

// Dfn attributes
type DfnAttributes struct {
    GlobalAttributes
}

// Div attributes
type DivAttributes struct {
    GlobalAttributes
}

// Dl attributes
type DlAttributes struct {
    GlobalAttributes
}

// Dt attributes
type DtAttributes struct {
    GlobalAttributes
}

// Em attributes
type EmAttributes struct {
    GlobalAttributes
}

// Figcaption attributes
type FigcaptionAttributes struct {
    GlobalAttributes
}

// Figure attributes
type FigureAttributes struct {
    GlobalAttributes
}

// Footer attributes
type FooterAttributes struct {
    GlobalAttributes
}

// Header attributes
type HeaderAttributes struct {
    GlobalAttributes
}

// H1, H2, H3, H4, H5, H6 attributes
type HeadingAttributes struct {
    GlobalAttributes
}

// Hr attributes
type HrAttributes struct {
    GlobalAttributes
}

// Html attributes
type HtmlAttributes struct {
    GlobalAttributes
    Manifest          string             `json:"manifest,omitempty"`
}

// I attributes
type IAttributes struct {
    GlobalAttributes
}

// Ins attributes
type InsAttributes struct {
    GlobalAttributes
    Cite              string             `json:"cite,omitempty"`
    DateTime          string             `json:"datetime,omitempty"`
}

// Kbd attributes
type KbdAttributes struct {
    GlobalAttributes
}

// Main attributes
type MainAttributes struct {
    GlobalAttributes
}

// Mark attributes
type MarkAttributes struct {
    GlobalAttributes
}

// Nav attributes
type NavAttributes struct {
    GlobalAttributes
}

// Noscript attributes
type NoscriptAttributes struct {
    GlobalAttributes
}

// P attributes
type PAttributes struct {
    GlobalAttributes
}

// Pre attributes
type PreAttributes struct {
    GlobalAttributes
}

// Ruby attributes
type RubyAttributes struct {
    GlobalAttributes
}

// S attributes
type SAttributes struct {
    GlobalAttributes
}

// Samp attributes
type SampAttributes struct {
    GlobalAttributes
}

// Section attributes
type SectionAttributes struct {
    GlobalAttributes
}

// Small attributes
type SmallAttributes struct {
    GlobalAttributes
}

// Span attributes
type SpanAttributes struct {
    GlobalAttributes
}

// Strong attributes
type StrongAttributes struct {
    GlobalAttributes
}

// Sub attributes
type SubAttributes struct {
    GlobalAttributes
}

// Summary attributes
type SummaryAttributes struct {
    GlobalAttributes
}

// Sup attributes
type SupAttributes struct {
    GlobalAttributes
}

// Template attributes
type TemplateAttributes struct {
    GlobalAttributes
}

// Tbody attributes
type TbodyAttributes struct {
    GlobalAttributes
}

// Td attributes
type TdAttributes struct {
    GlobalAttributes
    ColSpan           int                `json:"colspan,omitempty"`
    Headers           string             `json:"headers,omitempty"`
    RowSpan           int                `json:"rowspan,omitempty"`
}

// Tfoot attributes
type TfootAttributes struct {
    GlobalAttributes
}

// Th attributes
type ThAttributes struct {
    GlobalAttributes
    ColSpan           int                `json:"colspan,omitempty"`
    Headers           string             `json:"headers,omitempty"`
    RowSpan           int                `json:"rowspan,omitempty"`
    Scope             string             `json:"scope,omitempty"`
}

// Thead attributes
type TheadAttributes struct {
    GlobalAttributes
}

// Tr attributes
type TrAttributes struct {
    GlobalAttributes
}

// U attributes
type UAttributes struct {
    GlobalAttributes
}

// Var attributes
type VarAttributes struct {
    GlobalAttributes
}

// Wbr attributes
type WbrAttributes struct {
    GlobalAttributes
}

