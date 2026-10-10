package parser

import (
	"rpp-parser/pkg/lexer"
	"rpp-parser/pkg/token"
	"testing"
)

func TestCursorStatement(t *testing.T) {
	test := Project{
		Cursor: 1234,
	}

	input := "\n  CURSOR 1234\n"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	proj := &Project{}

	p.ParseProject(proj)

	if proj == nil {
		t.Fatal("failed to parse input. expected project got nil")
	}

	if proj.Cursor != test.Cursor {
		t.Fatalf("failed to currectly parse cursor statement expected %v got %v", test.Cursor, proj.Cursor)
	}

}

func TestMarkerProperties(t *testing.T) {
	tests := []Marker{
		{
			ID:       1,
			Position: 0,
			Name:     "",
			IsRegion: false,
			Color:    0,
			one:      1,
			b:        "B",
			GUID:     "5B24D1B6-3896-B803-0145-386DA9EF6B5B",
			IsEnd:    false,
		},
		{
			ID:       1,
			Position: 0,
			Name:     "You Are Good - Bethel Music - A - 144 BPM",
			IsRegion: true,
			Color:    0,
			one:      1,
			b:        "B",
			GUID:     "5B24D1B6-3896-B803-0145-386DA9EF6B5B",
			IsEnd:    false,
		},
		{
			ID:       1,
			Position: 3.33333333333333,
			Name:     "",
			IsEnd:    true,
		},
		{
			ID:       2,
			Position: 4.4,
			Name:     "Начало новой секции",
			IsRegion: true,
			Color:    0,
			one:      1,
			b:        "B",
			GUID:     "C3384881-AE94-41E0-F22F-7E8CEFDE76D6",
			IsEnd:    false,
		},
	}
	input := `
  MARKER 1 0 "" 0 0 1 B {5B24D1B6-3896-B803-0145-386DA9EF6B5B} 0
  MARKER 1 0 "You Are Good - Bethel Music - A - 144 BPM" 1 0 1 B {C2284881-AE94-41E0-F22F-7E8CEFDE76D6} 0
  MARKER 1 3.33333333333333 "" 1
  MARKER 2 4.4 "Начало новой секции" 1 0 1 B {C3384881-AE94-41E0-F22F-7E8CEFDE76D6} 0
`
	l := lexer.NewLexer(input)
	p := NewParser(l)

	proj := Project{}

	p.ParseProject(&proj)

	if len(proj.Markers) < 3 {
		t.Fatalf("Project.Properties does not contain 3 properties got=%d", len(proj.Markers))
	}

	for i, tt := range tests {
		m := proj.Markers[i]

		if m.ID != tt.ID {
			t.Fatalf("item[%d] - ID test failed. expected %v got %v", i, tt.ID, m.ID)
		}

		if m.Position != tt.Position {
			t.Fatalf("item[%d] - position test failed. expected %v got %v", i, tt.Position, m.Position)
		}

		if m.Name != tt.Name {
			t.Fatalf("item[%d] - name test failed. expected %v got %v", i, tt.Name, m.Name)
		}

		if m.IsRegion != tt.IsRegion {
			t.Fatalf("item[%d] - IsRegion test failed. expected %v got %v", i, tt.IsRegion, m.IsEnd)
		}

		if m.Color != tt.Color {
			t.Fatalf("item[%d] - color test failed. expected %v got %v", i, tt.Color, m.Color)
		}

	}

}

func TestParseItem(t *testing.T) {
	input := `
    <ITEM
      POSITION 0
      SNAPOFFS 0
      LENGTH 276.66666666666669
      LOOP 0
      ALLTAKES 0
      FADEIN 1 0 0 1 0 0 0
      FADEOUT 1 0 0 1 0 0 0
      MUTE 0 0
      SEL 1
      IGUID {B746E91F-4F81-7738-39B0-119903062AC2}
      IID 3594
      NAME "You Are Good - Bethel Music - A - 144 BPM.rpp"
      VOLPAN 1 0 1 -1
      SOFFS 0
      PLAYRATE 1 1 0 -1 0 0.0025
      CHANMODE 0
      GUID {DDB1FDB9-C10B-8C64-06FD-11902911C144}
      <SOURCE RPP_PROJECT
        FILE "/home/butbkadrug/Music/REAPER/You Are Good - Bethel Music - A - 144 BPM/You Are Good - Bethel Music - A - 144 BPM.rpp"
      >
    >
    `

	testItem := Item{
		Position: 0,
		Snapoffs: 0,
		Length:   276.66666666666669,
		Loop:     false,
		Alltakes: false,
		Fadein: struct {
			Curve      int
			FadeinTime float64
			OtherTime  float64
		}{1, 0, 0},
		Fadeout: struct {
			Curve      int
			FadeinTime float64
			OtherTime  float64
		}{1, 0, 0},
		Mute:  false,
		Sel:   true,
		IGUID: "B746E91F-4F81-7738-39B0-119903062AC2",
		IID:   3594,
		Name:  "You Are Good - Bethel Music - A - 144 BPM.rpp",
		Volpan: struct {
			Trim   float64
			Pan    float64
			Volume float64
			PanLaw float64
		}{
			1, 0, 1, -1},
		Soffs: 0,
		Playrate: struct {
			Playrate      float64
			PreservePitch bool
			PitchAdjust   float64
			PitchMode     int
		}{
			1, true, 0, -1},
		Chanmode: 0,
		GUID:     "DDB1FDB9-C10B-8C64-06FD-11902911C144",
		SourceSection: struct {
			Length   float64
			Mode     int
			Startpos float64
			Overlap  float64
			Source   struct {
				Format string
				File   string
			}
		}{
			0, 0, 0, 0, struct {
				Format string
				File   string
			}{
				"RPP_PROJECT", "/home/butbkadrug/Music/REAPER/You Are Good - Bethel Music - A - 144 BPM/You Are Good - Bethel Music - A - 144 BPM.rpp"}},
	}

	l := lexer.NewLexer(input)
	p := NewParser(l)

	item := p.parseItem()

	if item == nil {
		t.Fatalf("return item is nil. expected an Item")
	}

	if item.Name != testItem.Name {
		t.Fatalf("wrong item name expected %v got %v", testItem, item)
	}

	if item.Position != testItem.Position {
		t.Fatalf("wrong item position. expected %v got %v", testItem, item)
	}

	if item.Length != testItem.Length {
		t.Fatalf("item parsing failed. Expected item length %v got %v instead", testItem.Length, item.Length)
	}

	if item.SourceSection.Source.Format != testItem.SourceSection.Source.Format {
		t.Fatalf("source format check failed. expected %v got %v", testItem.SourceSection.Source.Format, item.SourceSection.Source.Format)
	}

	if item.SourceSection.Source.File != testItem.SourceSection.Source.File {
		t.Fatalf("source file check failed. expected %v got %v", testItem.SourceSection.Source.File, item.SourceSection.Source.File)
	}

	if p.curToken.Type != token.NEWLINE {
		t.Fatalf("item parsing failed. expected to be left with %T type of token got %T instead", token.NEWLINE, p.curToken.Type)
	}
}

func TestParseTrack(t *testing.T) {
	test := Track{}

	test.Name = "Click Track"
	test.TrackID = "5D232EF5-43E7-8751-3F38-4641E4BBF816"
	test.Items = append(test.Items,
		Item{
			Position: 0,
			Length:   276.66666666666669,
			Name:     "You Are Good - Bethel Music - A - 144 BPM.rpp",
			SourceSection: struct {
				Length   float64
				Mode     int
				Startpos float64
				Overlap  float64
				Source   struct {
					Format string
					File   string
				}
			}{
				Source: struct {
					Format string
					File   string
				}{
					Format: "RPP_PROJECT",
					File:   "/home/butbkadrug/Music/REAPER/You Are Good - Bethel Music - A - 144 BPM/You Are Good - Bethel Music - A - 144 BPM.rpp",
				},
			},
		},
		Item{
			Position: 1023.9898409475572,
			Length:   337.01015905244276,
			Name:     "Promises(VIKTAR Version) - Maverick City - Bb - 86 BPM.rpp",
			SourceSection: struct {
				Length   float64
				Mode     int
				Startpos float64
				Overlap  float64
				Source   struct {
					Format string
					File   string
				}
			}{
				Source: struct {
					Format string
					File   string
				}{
					Format: "RPP_PROJECT",
					File:   "/home/butbkadrug/Music/REAPER/Promises - Maverick City - Bb - 86 BPM/Promises(VIKTAR Version) - Maverick City - Bb - 86 BPM.rpp",
				},
			},
		},
		Item{
			Position: 2048,
			Length:   169.75609977324257,
			Name:     "Idi Vpered - Deti Boga - A - 123 BPM.RPP",
			SourceSection: struct {
				Length   float64
				Mode     int
				Startpos float64
				Overlap  float64
				Source   struct {
					Format string
					File   string
				}
			}{
				Source: struct {
					Format string
					File   string
				}{
					Format: "RPP_PROJECT",
					File:   "/home/butbkadrug/Music/REAPER/Idi Vpered - Deti Boga - A - 123 BPM/Idi Vpered - Deti Boga - A - 123 BPM.RPP",
				},
			},
		},
		Item{

			Position: 3072,
			Length:   290,
			Name:     "At The Cross (Love Ran Red) - Chris Tomlin - Ab - 73 BPM.rpp",
			SourceSection: struct {
				Length   float64
				Mode     int
				Startpos float64
				Overlap  float64
				Source   struct {
					Format string
					File   string
				}
			}{
				Source: struct {
					Format string
					File   string
				}{
					Format: "RPP_PROJECT",
					File:   "/home/butbkadrug/Music/REAPER/At The Cross (Love Ran Red) - Chris Tomlin - Ab - 73 BPM/At The Cross (Love Ran Red) - Chris Tomlin - Ab - 73 BPM.rpp",
				},
			},
		},
		Item{
			Position: 4096,
			Length:   205.50003405244297,
			Name:     "What The Lord Has Done in Me.rpp",
			SourceSection: struct {
				Length   float64
				Mode     int
				Startpos float64
				Overlap  float64
				Source   struct {
					Format string
					File   string
				}
			}{
				Source: struct {
					Format string
					File   string
				}{
					Format: "RPP_PROJECT",
					File:   "/home/butbkadrug/Music/REAPER/What The Lord Has Done in Me/What The Lord Has Done in Me.rpp",
				},
			},
		},
		Item{

			Position: 5120,
			Length:   284.8,
			Name:     "Lord, I Need You - Matt Maher - Bb - 75 BPM.rpp",
			SourceSection: struct {
				Length   float64
				Mode     int
				Startpos float64
				Overlap  float64
				Source   struct {
					Format string
					File   string
				}
			}{
				Source: struct {
					Format string
					File   string
				}{
					Format: "VIDEO",
					File:   "/home/butbkadrug/Downloads/m4a/Tremble - Drums.m4a",
				},
			},
		},
	)

	input := `<TRACK {5D232EF5-43E7-8751-3F38-4641E4BBF816}
    NAME "Click Track"
    PEAKCOL 16576
    BEAT -1
    AUTOMODE 0
    VOLPAN 0.12888250166462 0 -1 -1 1
    MUTESOLO 0 0 0
    IPHASE 0
    PLAYOFFS 0 1
    ISBUS 0 0
    BUSCOMP 0 0 0 0 0
    SHOWINMIX 1 0.6667 0.5 1 0.5 0 0 0
    FIXEDLANES 9 0 0
    SEL 1
    REC 0 0 1 0 0 0 0 0
    VU 2
    TRACKHEIGHT 0 0 0 0 0 0 0
    INQ 0 0 0 0.5 100 0 0 100
    NCHAN 4
    FX 1
    TRACKID {5D232EF5-43E7-8751-3F38-4641E4BBF816}
    PERF 0
    MIDIOUT -1
    MAINSEND 1 0
    HWOUT 1032 0 1 0 0 0 0 -1:U -1
    <ITEM
      POSITION 0
      SNAPOFFS 0
      LENGTH 276.66666666666669
      LOOP 0
      ALLTAKES 0
      FADEIN 1 0 0 1 0 0 0
      FADEOUT 1 0 0 1 0 0 0
      MUTE 0 0
      SEL 1
      IGUID {B746E91F-4F81-7738-39B0-119903062AC2}
      IID 3594
      NAME "You Are Good - Bethel Music - A - 144 BPM.rpp"
      VOLPAN 1 0 1 -1
      SOFFS 0
      PLAYRATE 1 1 0 -1 0 0.0025
      CHANMODE 0
      GUID {DDB1FDB9-C10B-8C64-06FD-11902911C144}
      <SOURCE RPP_PROJECT
        FILE "/home/butbkadrug/Music/REAPER/You Are Good - Bethel Music - A - 144 BPM/You Are Good - Bethel Music - A - 144 BPM.rpp"
      >
    >
    <ITEM
      POSITION 1023.9898409475572
      SNAPOFFS 0
      LENGTH 337.01015905244276
      LOOP 0
      ALLTAKES 0
      FADEIN 1 0 0 1 0 0 0
      FADEOUT 1 0 0 1 0 0 0
      MUTE 0 0
      SEL 1
      IGUID {5B3CF868-59A8-056B-13DD-A0DADC2F4D0C}
      IID 3595
      NAME "Promises(VIKTAR Version) - Maverick City - Bb - 86 BPM.rpp"
      VOLPAN 1 0 1 -1
      SOFFS 0
      PLAYRATE 1 1 0 -1 0 0.0025
      CHANMODE 0
      GUID {3E6B7F3B-C446-A7F1-6C1A-9549B0B3344B}
      <SOURCE RPP_PROJECT
        FILE "/home/butbkadrug/Music/REAPER/Promises - Maverick City - Bb - 86 BPM/Promises(VIKTAR Version) - Maverick City - Bb - 86 BPM.rpp"
      >
    >
    <ITEM
      POSITION 2048
      SNAPOFFS 0
      LENGTH 169.75609977324257
      LOOP 0
      ALLTAKES 0
      FADEIN 1 0 0 1 0 0 0
      FADEOUT 1 0 0 1 0 0 0
      MUTE 0 0
      SEL 1
      IGUID {E3DD80AC-B20A-A2A3-D1C1-2DB21D8437BE}
      IID 84
      NAME "Idi Vpered - Deti Boga - A - 123 BPM.RPP"
      VOLPAN 1 0 1 -1
      SOFFS 0
      PLAYRATE 1 1 0 -1 0 0.0025
      CHANMODE 0
      GUID {BB457C93-D46F-75EC-A3EC-A44E9D34C26C}
      <SOURCE RPP_PROJECT
        FILE "/home/butbkadrug/Music/REAPER/Idi Vpered - Deti Boga - A - 123 BPM/Idi Vpered - Deti Boga - A - 123 BPM.RPP"
      >
    >
    <ITEM
      POSITION 3072
      SNAPOFFS 0
      LENGTH 290
      LOOP 0
      ALLTAKES 0
      FADEIN 1 0 0 1 0 0 0
      FADEOUT 1 0 0 1 0 0 0
      MUTE 0 0
      SEL 1
      IGUID {E0A67297-11D1-9995-91CE-74C171CEA506}
      IID 3599
      NAME "At The Cross (Love Ran Red) - Chris Tomlin - Ab - 73 BPM.rpp"
      VOLPAN 1 0 1 -1
      SOFFS 0
      PLAYRATE 1 1 0 -1 0 0.0025
      CHANMODE 0
      GUID {76797C93-F80E-04F2-767D-61753B55A9CA}
      <SOURCE RPP_PROJECT
        FILE "/home/butbkadrug/Music/REAPER/At The Cross (Love Ran Red) - Chris Tomlin - Ab - 73 BPM/At The Cross (Love Ran Red) - Chris Tomlin - Ab - 73 BPM.rpp"
      >
    >
    <ITEM
      POSITION 4096
      SNAPOFFS 0
      LENGTH 205.50003405244297
      LOOP 0
      ALLTAKES 0
      FADEIN 1 0 0 1 0 0 0
      FADEOUT 1 0 0 1 0 0 0
      MUTE 0 0
      SEL 1
      IGUID {46CD352A-4DD9-C99A-331F-30F72297468C}
      IID 3598
      NAME "What The Lord Has Done in Me.rpp"
      VOLPAN 1 0 1 -1
      SOFFS 0
      PLAYRATE 1 1 0 -1 0 0.0025
      CHANMODE 0
      GUID {CE2CB567-0899-9D84-16EB-0BE1DA296ED0}
      <SOURCE RPP_PROJECT
        FILE "/home/butbkadrug/Music/REAPER/What The Lord Has Done in Me/What The Lord Has Done in Me.rpp"
      >
    >
    <ITEM
      POSITION 5120
      SNAPOFFS 0
      LENGTH 284.8
      LOOP 1
      ALLTAKES 0
      FADEIN 1 0 0 1 0 0 0
      FADEOUT 1 0 0 1 0 0 0
      MUTE 0 0
      SEL 1
      IGUID {BADE797E-F708-EFAC-6A84-842309181AE4}
      IID 183
      NAME "Lord, I Need You - Matt Maher - Bb - 75 BPM.rpp"
      VOLPAN 1 0 1 -1
      SOFFS 0
      PLAYRATE 1 1 0 -1 0 0.0025
      CHANMODE 0
      GUID {A53F3456-11BE-01F5-AF57-ECBB53A36ED1}
      <SOURCE VIDEO
        FLIP 0 1
        FILE "/home/butbkadrug/Downloads/m4a/Tremble - Drums.m4a"
      >
    >
  >
    `

	l := lexer.NewLexer(input)
	p := NewParser(l)

	track := p.parseTrackStatement()

	if track.Name != test.Name {
		t.Fatalf("track parsing test failed. expected track name %s got %s", test.Name, track.Name)
	}

	if track.TrackID != test.TrackID {
		t.Fatalf("track parsing test failed. expected track id %s got %s", test.TrackID, track.TrackID)
	}

	if len(track.Items) < 6 {
		t.Fatalf("track parsing test failed. expected 6 items got = %d", len(track.Items))
	}

	for i, tt := range test.Items {
		item := track.Items[i]
		if item.Position != tt.Position {
			t.Fatalf("track.items[%d] parsing test failed expected position %v got %v", i, tt.Position, item.Position)
		}

		if item.Length != tt.Length {
			t.Fatalf("track.items[%d] parsing test failed expected length %v got %v", i, tt.Length, item.Length)
		}

		if item.Name != tt.Name {
			t.Fatalf("track.items[%d] parsing test failed expected name %v got %v", i, tt.Name, item.Name)
		}

		if item.SourceSection.Source.File != tt.SourceSection.Source.File {
			t.Fatalf("track.items[%d] parsing test failed source file expected %v got %v", i, tt.SourceSection.Source.File, item.SourceSection.Source.File)
		}
		if item.SourceSection.Source.Format != tt.SourceSection.Source.Format {
			t.Fatalf("track.items[%d] parsing test failed source format expected %v got %v", i, tt.SourceSection.Source.Format, item.SourceSection.Source.Format)
		}
	}

}
