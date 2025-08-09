package parser

import (
	"log"
	"log/slog"
	"rpp-parser/pkg/lexer"
	"rpp-parser/pkg/token"
	"strconv"
)

type Project struct {
	// Project title
	Name string

	// Project author
	Author string

	//Project Notes
	Notes []string

	//Cursor position
	Cursor float64

	// Path to a project file
	Path string

	// Information about the project. Build information and under which system it has been saved
	Version struct {
		// Unknown value
		Status  float64
		Version string
		Build   int
	}

	// Markers and Regions parsed and stored as found in reaper.rpp file
	Markers []Marker
	// Tracks parsed and stored as found in reaper.rpp file
	Tracks []Track
}

type Marker struct {
	// Markers ID potentially not uniqe between regions and markers
	ID int

	// Position on the timelene in seconds
	Position float64

	// Description text of the marker/region. Text that seen in GUI
	Name string

	// Boolean flag tha let's us know wether this marker is a region
	// thus will be followed by another marker that represents the end
	// of the region
	IsRegion bool

	// Native color. defaults to 0.
	Color int

	// I have ho idea what's this value
	one int

	// I have no idea what's this value ether
	b string
	// Unique identifired of a  marker
	GUID string

	// Bool flag that lets us know where this marker represents the end of the region
	IsEnd bool
}

func newMarker() *Marker {
	return &Marker{
		IsRegion: false,
		IsEnd:    false,
	}
}

func (m *Marker) TokenLiteral() string { return "MARKER" }

type Track struct {
	// Unique identifier of a track
	GUID string

	// The Name of the track
	Name string

	Peakcol     int
	Beat        int
	Automode    int
	VolPan      []int
	MuteSolo    []int
	Iphase      int
	PLAYOFFS    []int
	IsBus       []int
	BusComp     []int
	ShowInMix   []float64
	FixedLanes  []int
	Sel         bool
	Rec         []int
	VU          int
	TrackHeight []int
	INQ         []float64
	Nchan       int
	FX          int
	TrackID     string
	Perf        int
	AuxRecv     []any
	MIDIout     int
	MainSend    []int
	HWOUT       []any
	Items       []Item
}

func (t *Track) TokenLiteral() string { return "TRACK" }

type Item struct {

	// Position on the timeline, in seconds
	Position float64

	// Snap offset, in seconds
	Snapoffs float64

	// Item length, in seconds
	Length float64

	// Loop source? int (bool)
	Loop bool

	// Play all takes? int (bool)
	Alltakes bool

	// Item colour, int , Not present if not set.
	Color int

	// Item timebase, int, Not present if not set. -1 = use track timebase
	Beat int

	// Is the item selected? int (bool)
	Sel bool

	// Fade in settings
	//  field 1, int, fade in curve type (0 - linear, etc.)
	//  field 2, float, fade in time (in seconds)
	//  field 3, float, ??
	Fadein struct {
		Curve      int
		FadeinTime float64
		OtherTime  float64
	}

	// Fade out settings
	//  field 1, int, fade in curve type (0 - linear, etc.)
	//  field 2, float, fade in time (in seconds)
	//  field 3, float, ??
	Fadeout struct {
		Curve      int
		FadeinTime float64
		OtherTime  float64
	}

	// Is the item muted? int (bool)
	Mute bool

	// ??, int, may not be present
	Fadeflag int

	// A number of the group the item belongs to, if grouped
	Group int

	// Item id. Leave it alone
	// --- From this point, values mostly relate to the first take
	//     except for field 1 of VOLPAN which is the item trim
	//
	IGUID string
	// Deprecated since build 6.54
	// Item ordinal number governing the order of overlapping items
	// displayed in lanes and especially those with identical start
	// and end coordinates, when
	// 'Show overlapping media items in lanes (when room)' option
	// is enabled. The greater the IID the lower the item's lane is.
	// With the said option disabled or lanes fully collapsed
	// the visible/outermost item is the latest one on the track
	// out of all items it's overlapping, regardless of its IID
	IID int

	// Notes block
	Notes []string

	// Name of the image file loaded in the notes, full path if located outside
	// of the resource directory
	Resourcefn string

	//Unknown option
	Imgresourceflags int

	// First take name
	Name string

	// Volume and pan settings
	// field 1, float, item trim
	//   1.00 = 0 dB
	//   16.00 = +24 dB
	//   0.50 = -6 dB
	// field 2, float, first take pan, -1.00 = left, 1.00 = right
	// field 3, float, first take volume (other takes in TAKEVOLPAN below),
	// if negative 'Invert phase' setting is enabled
	// field 4, float, first take pan law
	Volpan struct {
		Trim   float64
		Pan    float64
		Volume float64
		PanLaw float64
	}

	// First take slip offset (in seconds)
	Soffs float64

	// First take playrate settings
	//  field 1, float, playrate
	//  field 2, int (bool), preserve pitch while changing rate
	//  field 3, float, pitch adjust, in semitones.cents
	//  field 4, int, pitch shifting/time stretch mode and preset
	//    -1 - project default
	//    0-2 - Sound Touch (3 presets)
	//    65536-65567 - Dirac LE (32 presets)
	//    131072-131119 - Low quality windowed (48 presets)
	//    196608-196639 - élastique Pro (32 presets)
	//    262144-262147 - élastique Efficient (4 presets)
	//    327680-327683 - élastique SOLOIST (4 presets)
	//    393216-393247 - élastique 2.1 Pro (32 presets)
	//    458752-458755 - élastique 2.1 Efficient (4 presets)
	//    524288-524291 - élastique 2.1 SOLOIST (4 presets)
	Playrate struct {
		Playrate      float64
		PreservePitch bool
		PitchAdjust   float64
		PitchMode     int
	}

	// First take channel mode, int
	//   0 = normal, 1 = reverse stereo, 2 = mono (downmix)
	//   3 = mono (left), 4 = mono (right)
	Chanmode int

	// First take id. Leave it alone
	GUID string

	// Source section start (only present when changed in Item properties)
	SourceSection struct {
		// Section length
		Length float64

		// Mode, 2 = Section and Reverse, 3 = Reverse
		Mode int

		// Section start offset (in seconds)
		Startpos float64
		// Section crossfade (in seconds)
		Overlap float64
		// Source details. Appears alone if SOURCE SECTION not set

		Source struct {
			//File format
			Format string
			// Source file path.
			File string
		}
	}

	// Take marker data
	// field 1 - float, position within take in sec
	// field 2 - string, name, without quotes if consists of only one
	// character or word
	// field 3 - integer, OS native color code
	// can be converted to RGB with: r, g, b = reaper.ColorFromNative(code)
	TKM struct {
		Position float64
		Name     string
		Color    int
	}
}

func (i *Item) TokenLiteral() string { return "ITEM" }

type Parser struct {
	l      *lexer.Lexer
	Errors []string

	curToken  token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		Errors: []string{},
	}

	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseMarkerStatement() *Marker {
	m := newMarker()

	//NOTE: Parsing ID

	if !p.expectPeek(token.INT) {
		log.Fatal("error parsing marker id. expected int, got ", p.peekToken.Type)
	}

	id, err := strconv.Atoi(p.curToken.Value)

	if err != nil {
		log.Fatal("error parsing marker: ", err)
	}

	m.ID = id

	//NOTE: Parsing position

	m.Position = p.parsePositionStatement()

	//NOTE: Parsin Name

	m.Name = p.parseNameStatement()

	//NOTE: Parsing isRegion or Ending

	if !p.expectPeek(token.INT) {
		log.Fatal("error parsing marker type exprected int got ", p.peekToken.Type)
	}

	valueInt, err := strconv.Atoi(p.curToken.Value)

	isRegion := valueInt != 0 && valueInt != 8

	if err != nil {
		log.Fatal("error parsing marker type ", err)
	}

	//NOTE: if marker signifies the end of the region we stop here and return
	if p.expectPeek(token.NEWLINE) {
		m.IsEnd = true

		return m
	}

	//NOTE: otherwise we will continue process data
	m.IsRegion = isRegion

	//NOTE: Color

	if !p.expectPeek(token.INT) {
		log.Fatal("error parsing marker color. expected int got ", p.peekToken.Type)
	}

	color, err := strconv.Atoi(p.curToken.Value)

	if err != nil {
		log.Fatal("error parsing marker color ", err)
	}

	m.Color = color

	//NOTE: Unknown field

	if !p.expectPeek(token.INT) {
		log.Fatal("error parsing marker unknown field. expected int got ", p.peekToken.Type)
	}

	one, err := strconv.Atoi(p.curToken.Value)

	if err != nil {
		log.Fatal("error parsing marker unknown field: ", err)
	}

	m.one = one

	//NOTE: B field

	if !p.expectPeek(token.STRING) {
		log.Fatal("error parsing marker B field. expected ident got ", p.peekToken.Type)
	}

	m.b = p.curToken.Value

	//NOTE: GUID

	if !p.expectPeek(token.STRING) {
		log.Fatal("error parsing marker guid expected ident got ", p.peekToken.Type)
	}

	m.GUID = p.curToken.Value

	//NOTE: End

	if !p.expectPeek(token.INT) {
		log.Fatal("error parsing marker isEnd. expected int got ", p.peekToken.Type)
	}

	end, err := strconv.ParseBool(p.curToken.Value)

	if err != nil {
		log.Fatal("error parsing marker isEnd: ", err)
	}

	m.IsEnd = end

	if !p.expectPeek(token.NEWLINE) {
		log.Fatal("error parsing marker. expected new line got data:", p.peekToken.Type, p.peekToken.Value)
	}

	return m
}

func (p *Parser) parseFloat(stringValue, errMsg string) float64 {
	value, err := strconv.ParseFloat(stringValue, 64)

	if err != nil {
		log.Fatal(errMsg, err)
	}

	return value
}

func (p *Parser) parseInt(stringValue, errMsg string) int {
	value, err := strconv.Atoi(stringValue)

	if err != nil {
		log.Fatal(errMsg, err)
	}

	return value
}

func (p *Parser) unexpectedPeekError(itemName, expectedType string) {

	log.Fatalf("error parsing project %s. expected %s got %s", itemName, expectedType, p.peekToken.Type)

}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	}

	return false
}

func (p *Parser) parseItem() *Item {

	item := &Item{}

	slog.Debug("entered parseItem", "tok", p.curToken, "peek", p.peekToken)

	for !p.curTokenIs(token.GT) {
		slog.Debug("parse item log", "tok", p.curToken)
		switch p.curToken.Type {
		case token.POSITION:
			item.Position = p.parsePositionStatement()
		case token.LENGTH:
			item.Length = p.parseLengthStatement()
		case token.NAME:
			item.Name = p.parseNameStatement()
		case token.SOURCE:
			//BUG: click source not gonna parse currectly
			//BUG: midi source gives me some minor trouble too
			if !p.expectPeek(token.STRING) {
				p.unexpectedPeekError("item source type", token.STRING)
			}

			if p.curToken.Value == "CLICK" || p.curToken.Value == "MIDI" {
				return nil
			}

			item.SourceSection.Source.Format = p.curToken.Value

			if !p.expectPeek(token.NEWLINE) {
				p.unexpectedPeekError("item source", token.NEWLINE)
			}

			for !p.peekTokenIs(token.FILE) && !p.peekTokenIs(token.GT) {
				slog.Debug("skipping until file section", "tok", p.curToken, "peek", p.peekToken)
				p.NextToken()
			}

			if !p.expectPeek(token.FILE) && !p.expectPeek(token.GT) {
				p.unexpectedPeekError("item File", token.FILE)
			}

			if p.expectPeek(token.GT) {
				slog.Debug("shuld return if source not valid", "tok", p.curToken)
				return item
			}

			if !p.expectPeek(token.STRING) {
				p.unexpectedPeekError("item file", token.STRING)
			}

			item.SourceSection.Source.File = p.curToken.Value

		}
		p.NextToken()
	}

	// consuming '>' to leave parser with a \n token
	p.NextToken()

	return item
}

func (p *Parser) skipClosure(closingIdent token.TokenType) {
	for p.curToken.Type != closingIdent && p.curToken.Type != token.EOF {
		p.NextToken()
	}

	p.NextToken()
}

func (p *Parser) parseTrackStatement() *Track {

	track := &Track{}

	for !p.peekTokenIs(token.GT) && !p.peekTokenIs(token.EOF) {
		switch p.curToken.Type {
		case token.LT:
			slog.Debug("LT", "token", p.curToken, "peek", p.peekToken)
			if !p.peekTokenIs(token.ITEM) {
				p.skipClosure(token.GT)
			}

		case token.NAME:
			slog.Debug("Name", "token", p.curToken, "peek", p.peekToken)
			track.Name = p.parseNameStatement()
			slog.Debug("Name exit", "token", p.curToken, "peek", p.peekToken)
		case token.TRACKID:
			slog.Debug("ID ", "token", p.curToken, "peek", p.peekToken)
			if !p.expectPeek(token.STRING) {
				p.unexpectedPeekError("track ID", token.STRING)
			}
			track.TrackID = p.curToken.Value
		case token.ITEM:
			slog.Debug("case item ", "curent token", p.curToken)
			i := p.parseItem()
			if i != nil {
				track.Items = append(track.Items, *i)
			}
		default:
			slog.Debug("parseTrackStatement()default", "token", p.curToken)
		}
		p.NextToken()
	}

	return track
}

func (p *Parser) parseNameStatement() string {
	p.NextToken()

	return p.curToken.Value
}

func (p *Parser) parsePositionStatement() float64 {
	if !p.expectPeek(token.INT) && !p.expectPeek(token.FLOAT) {
		log.Fatal("error parsing marker position. expected float, got ", p.peekToken.Type)
	}

	return p.parseFloat(p.curToken.Value, "failed to parse position statement")
}

func (p *Parser) parseLengthStatement() float64 {

	if !p.expectPeek(token.FLOAT) && !p.expectPeek(token.INT) {
		p.unexpectedPeekError("item length", token.FLOAT)
	}

	return p.parseFloat(p.curToken.Value, "error parsing item length")
}

func (p *Parser) parseCursorStatement() float64 {
	if !p.expectPeek(token.FLOAT) && !p.expectPeek(token.INT) {
		p.unexpectedPeekError("cursor", token.FLOAT)
	}

	return p.parseFloat(p.curToken.Value, "error parsing project cursor")
}

func (p *Parser) ParseProject(project *Project) {

	for p.curToken.Type != token.EOF {
		slog.Debug("new token", "token", p.curToken, "peek", p.peekToken)
		switch p.curToken.Type {
		case token.MARKER:
			m := p.parseMarkerStatement()
			if m != nil {
				project.Markers = append(project.Markers, *m)
			}
		case token.TRACK:
			slog.Debug("track state")
			track := p.parseTrackStatement()
			if track != nil {
				project.Tracks = append(project.Tracks, *track)
			}
		case token.NAME:
			project.Name = p.parseNameStatement()
		case token.CURSOR:
			project.Cursor = p.parseCursorStatement()
		default:
			slog.Debug("parser: default case no op", "token", p.curToken, "peek", p.peekToken)
		}

		p.NextToken()
	}

	slog.Debug("parser: returning from parseProject()", "token", p.curToken, "peek", p.peekToken)

}
