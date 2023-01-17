package accelhint

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func Test001(t *testing.T) {
	original := []string{"O&ne", "Two"}
	expected := []string{"O&ne", "&Two"}
	hinted, count, err := Hinted(original)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	sanityCheck(hinted, t)
	if count != 2 {
		t.Errorf("expected 2 accelrated got %d", count)
	}
	expectedAccels := []rune{'n', 'T'}
	accels := Accelerators(hinted)
	if !slices.Equal(accels, expectedAccels) {
		t.Errorf("expected %v accels, got %v", expectedAccels, accels)
	}
	for i := 0; i < len(original); i++ {
		if hinted[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], hinted[i])
		}
	}
}

func Test002(t *testing.T) {
	original := []string{
		"Undo",
		"Redo",
		"Copy",
		"Cu&t",
		"Paste",
		"Find",
		"Find Again",
	}
	expected := []string{
		"&Undo",
		"&Redo",
		"&Copy",
		"Cu&t",
		"&Paste",
		"&Find",
		"Find &Again"}
	hinted, count, err := Hinted(original)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	sanityCheck(hinted, t)
	if count != 7 {
		t.Errorf("expected 7 accelrated got %d", count)
	}
	expectedAccels := []rune{'U', 'R', 'C', 't', 'P', 'F', 'A'}
	accels := Accelerators(hinted)
	if !slices.Equal(accels, expectedAccels) {
		t.Errorf("expected %v accels, got %v", expectedAccels, accels)
	}
	for i := 0; i < len(original); i++ {
		if hinted[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], hinted[i])
		}
	}
}

func Test003(t *testing.T) {
	original := []string{
		"Undo",
		"Redo",
		"Copy",
		"Cu&t",
		"Paste",
		"Find",
		"Find Again",
		"Find && Replace"}
	expected := []string{
		"&Undo",
		"&Redo",
		"&Copy",
		"Cu&t",
		"&Paste",
		"&Find",
		"Find &Again",
		"F&ind && Replace"}
	hinted, count, err := Hinted(original)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if count != 8 {
		t.Errorf("expected 8 accelrated got %d", count)
	}
	expectedAccels := []rune{'U', 'R', 'C', 't', 'P', 'F', 'A', 'i'}
	accels := Accelerators(hinted)
	if !slices.Equal(accels, expectedAccels) {
		t.Errorf("expected %v accels, got %v", expectedAccels, accels)
	}
	for i := 0; i < len(original); i++ {
		if hinted[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], hinted[i])
		}
	}
}

func Test004(t *testing.T) {
	originals := [][]string{
		{
			"&Bold",
			"Italic",
			"Underline",
			"No Super- or Sub-script",
			"Superscript",
			"Subscript",
			"Text Color",
			"Font",
			"No List",
			"Bullet List",
			"Numbered List",
			"Align Left",
			"Center",
			"Justify",
			"Align Right",
		},
		{
			"abc",
			"bca",
			"cab",
			"aab",
			"bbc",
			"cca",
			"cba",
			"bcb",
			"acc",
		},
		{
			"Lists",
			"New List",
			"Duplicate List",
			"Delete List",
			"Items",
			"New",
			"Rename",
			"Move &Up",
			"Move &Down",
			"Delete",
			"Configure",
			"About",
			"Quit",
		},
		{
			"Calculate",
			"Load...",
			"Save",
			"Save &As...",
			"Copy to Clipboard",
			"Set Options...",
			"Help",
			"About",
			"Quit",
		},
		{
			"Alphabet",
			"Predefined Accelerators",
			"Max. Recursion Depth",
			"Max. Candidates",
			"Automatic Timeout",
			"Prompt to Save Unsaved Changes on Exit",
			"Restore Workspace at Startup",
			"Load Last File at Startup",
			"Show Help at Startup",
		},
		{
			"Add",
			"Delete",
			"Edit",
		},
		{
			"Cu&t",
			"&Copy",
			"Paste",
			"Find && Replace",
			"Replace",
		},
		{
			"Basic",
			"Advanced",
			"Automatically Save on Exit",
			"Restore Workspace on Startup",
			"Remember Clipboard between Sessions",
		},
		{
			"Notebook",
			"Edit",
			"Clipboard",
			"Promotion",
			"Notes",
			"Calendar",
			"Calendar &View",
			"Calendar &Goto",
			"Note",
			"Clipboard",
			"Edit",
			"Search",
			"Style",
			"Font",
			"Alignment",
			"Category",
		},
		{
			"Type",
			"Title",
			"Forenames",
			"Surname",
			"Company",
			"Email",
			"Phone",
			"Mobile",
			"URL",
			"Find",
			"New",
			"Duplicate",
			"Update",
			"Delete",
			"Configure",
			"Statistics",
			"Export",
			"Open",
			"Save",
			"Quit",
		},
		{
			"Remaining Issues",
			"Type",
			"Title",
			"Forenames",
			"Surname",
			"Company",
			"Email",
			"Phone",
			"Mobile",
			"URL",
			"Find",
			"New",
			"Duplicate",
			"Update",
			"Delete",
			"Configure",
			"Statistics",
			"Export",
			"Open",
			"Save",
			"Quit",
		},
		{
			"Initial No",
			"Max Supportable No",
			"Initial Memory Size",
			"Mutation Rate",
			"Survival Rate",
			"Start",
			"Pa&use",
			"Zoom",
		},
		{
			"General",
			"Remember Geometry",
			"Start With Editor &Window",
			"Start With &Shell Window",
			"Automatically Hide the Find and Replace dialog",
			"Input prompt",
			"Output prompt",
			"&Python executable",
			"Working directory",
			"Show Window Information at Startup",
			"Max Tooltip Length",
			"Max Lines to Scan",
			"Highlighting",
			"Tab width",
			"Font",
			"On New File",
			"At Startup",
			"Python Documentation",
		},
		{
			"Undo",
			"Redo",
			"&Copy",
			"Cu&t",
			"&Paste",
			"Select &All",
			"Complete",
			"&Find",
			"Find &Next",
			"R&eplace Next",
			"Goto Line",
			"Goto &Matching",
			"Indent Region",
			"Unindent Region",
			"Comment Region",
			"Uncomment Region",
		},
		{
			"Add &Top-Level Term",
			"Add Sub-Term &of",
			"Add &Child-Term of",
			"&Add Sibling-Term of",
			"&Find",
			"&Insert Unicode Character",
			"&Bold",
			"Ita&lic",
			"&Underline",
			"Stri&ke-out",
			"Su&perscript",
			"&Subscript",
			"&No Superscript or Subscript",
			"&Monospace",
			"Sans S&erif",
			"Se&rif",
		},
		{
			"Toolbars",
			"Status Bar",
			"Sidebar",
			"Stop",
			"Reload",
			"Text Size",
			"Page Style",
			"Character Encoding",
			"Page Source",
			"Full Screen",
		},
		{
			"New",
			"Open...",
			"Recent Documents",
			"Wizards",
			"Close",
			"Save",
			"Save As...",
			"Save All",
			"Reload",
			"Versions...",
			"Export...",
			"Export as PDF...",
			"Send",
			"Properties...",
			"Digital Signatures...",
			"Templates",
			"Preview in Web Browser",
			"Page Preview",
			"Print...",
			"Printer Settings...",
			"Exit",
		},
		{
			"Undo",
			"Redo",
			"Cut",
			"Copy",
			"Paste",
			"Paste Special...",
			"Select Text",
			"Select All",
			"Changes",
			"Compare Documents...",
			"Find and Replace",
			"Navigator",
			"AutoText...",
			"Exchange Database...",
			"Fields...",
			"Footnote...",
			"Index Entry...",
			"Bibliography Entry",
			"Hyperlink",
			"Links...",
			"Plug-in",
			"ImageMap",
			"Object",
		},
		{
			"Print Layout",
			"Web Layout",
			"Toolbars",
			"Status Bar",
			"Input Method Status",
			"Ruler",
			"Text Boundaries",
			"Field Shadings",
			"Field Names",
			"Nonprinting Characters",
			"Hidden Paragraphs",
			"Data Sources",
			"Full Screen",
			"Zoom",
		},
		{
			"Manual Break",
			"Fields",
			"Special Character...",
			"Formatting Mark",
			"Section...",
			"Hyperlink",
			"Header",
			"Footer",
			"Footnote",
			"Caption",
			"Bookmark",
			"Cross-reference...",
			"Note...",
			"Script...",
			"Indexes and Tables",
			"Envelope...",
			"Frame...",
			"Table...",
			"Horizontal Ruler...",
			"Picture",
			"Movie and Sound",
			"Object",
			"Floating Frame",
			"File...",
		},
		{
			"Default Formatting",
			"Character...",
			"Paragraph...",
			"Bullets and Numbering...",
			"Page...",
			"Title Page",
			"Change Case",
			"Columns",
			"Sections...",
			"Styles and Formatting",
			"AutoFormat",
			"Anchor",
			"Wrap",
			"Alignment",
			"Arrange",
			"Flip",
			"Group",
			"Object",
			"Frame...",
			"Picture...",
		},
		{
			"Insert",
			"Delete",
			"Select",
			"Merge Cells",
			"Split Cells",
			"Protect Cells",
			"Merge Table",
			"Split Table",
			"AutoFormat...",
			"Autofit",
			"Heading rows repeat",
			"Convert",
			"Sort...",
			"Formula",
			"Number Format...",
			"Table Boundaries",
			"Table Properties...",
		},
		{
			"Spellcheck...",
			"Language",
			"Word Count",
			"AutoCorrect...",
			"Outline Numbering...",
			"Footnotes...",
			"Gallery",
			"Media Player",
			"Bibliography Database",
			"Mail Merge Wizard...",
			"Sort...",
			"Calculate",
			"Update",
			"Macros",
			"Extension Manager...",
			"XML Filter Settings...",
			"Customize...",
			"Options...",
		},
		{
			"&Undo",
			"&Redo",
			"Cu&t",
			"&Copy",
			"&Paste",
			"Paste Special...",
			"Select Text",
			"Select All",
			"Changes",
			"Compare Documents...",
			"Find and Replace",
			"Navigator",
			"AutoText...",
			"Exchange Database...",
			"Fields...",
			"Footnote...",
			"Index Entry...",
			"Bibliography Entry",
			"Hyperlink",
			"Links...",
			"Plug-in",
			"ImageMap",
			"Object",
		},
		{
			"&Fields",
			"Special Character...",
			"Formatting Mark",
			"Section...",
			"Hyperlink",
			"Header",
			"Footer",
			"Footnote",
			"Caption",
			"Bookmark",
			"Cross-reference...",
			"&Note...",
			"Script...",
			"Indexes and Tables",
			"Envelope...",
			"Frame...",
			"Table...",
			"Horizontal Ruler...",
			"Picture",
			"Movie and Sound",
			"Object",
			"Floating Frame",
		},
		{
			"Undo",
			"Redo",
			"Repeat",
			"Cu&t",
			"&Copy",
			"Paste",
			"Paste Special...",
			"Select All",
			"Changes",
			"Compare Document...",
			"Find and Replace...",
			"Navigator",
			"Headers and Footers...",
			"Fill",
			"Delete Contents...",
			"Delete Cells...",
			"Sheet",
			"Delete Manual Break",
			"Links...",
			"Plug-in",
			"ImageMap",
			"Object",
		},
		{
			"Undo",
			"Redo",
			"Repeat",
			"Cu&t",
			"&Copy",
			"Paste",
			"Paste Special...",
			"Select All",
			"Changes",
			"Compare Document...",
			"Find and Replace...",
			"Navigator",
			"Headers and Footers...",
			"Fill",
			"Delete Contents...",
			"Delete Cells...",
			"Sheet",
			"Delete Manual Break",
			"Links...",
			"Plug-in",
			"ImageMap",
			"Object",
		},
		{
			"Undo",
			"Redo",
			"&Copy",
			"Cu&t",
			"&Paste",
			"Select &All",
			"Complete",
			"&Find",
			"Find &Next",
			"R&eplace Next",
			"Goto Line",
			"Goto &Matching",
			"Indent Region",
			"Unindent Region",
			"Comment Region",
			"Uncomment Region",
		},
		{
			"Undo",
			"Redo",
			"Copy",
			"Cu&t",
			"Paste",
			"Find",
			"Find Again",
			"Find && Replace",
		},
		{"O&ne", "Two"},
		{
			"Undo",
			"Redo",
			"Copy",
			"Cu&t",
			"Paste",
			"Find",
			"Find Again",
		},
	}

	expecteds := [][]string{
		{ // 0
			"&Bold",
			"&Italic",
			"&Underline",
			"No Super- &or Sub-script",
			"Su&perscript",
			"&Subscript",
			"&Text Color",
			"&Font",
			"No &List",
			"Bull&et List",
			"&Numbered List",
			"&Align Left",
			"&Center",
			"&Justify",
			"Align &Right",
		},
		{ // 1
			"abc",
			"bca",
			"cab",
			"aab",
			"bbc",
			"cca",
			"&cba",
			"&bcb",
			"&acc",
		},
		{ // 2
			"&Lists",
			"&New List",
			"Du&plicate List",
			"D&elete List",
			"&Items",
			"Ne&w",
			"&Rename",
			"Move &Up",
			"Move &Down",
			"Dele&te",
			"&Configure",
			"&About",
			"&Quit",
		},
		{ // 3
			"&Calculate",
			"&Load...",
			"&Save",
			"Save &As...",
			"Copy &to Clipboard",
			"Set &Options...",
			"&Help",
			"A&bout",
			"&Quit",
		},
		{ // 4
			"&Alphabet",
			"&Predefined Accelerators",
			"&Max. Recursion Depth",
			"Max. &Candidates",
			"Automatic &Timeout",
			"Prompt to Save &Unsaved Changes on Exit",
			"&Restore Workspace at Startup",
			"&Load Last File at Startup",
			"&Show Help at Startup",
		},
		{ // 5
			"&Add",
			"&Delete",
			"&Edit",
		},
		{ // 6
			"Cu&t",
			"&Copy",
			"&Paste",
			"&Find && Replace",
			"&Replace",
		},
		{ // 7
			"&Basic",
			"&Advanced",
			"Automatically &Save on Exit",
			"Restore &Workspace on Startup",
			"&Remember Clipboard between Sessions",
		},
		{ // 8
			"Note&book",
			"&Edit",
			"C&lipboard",
			"&Promotion",
			"N&otes",
			"&Calendar",
			"Calendar &View",
			"Calendar &Goto",
			"&Note",
			"Cl&ipboard",
			"E&dit",
			"&Search",
			"St&yle",
			"&Font",
			"&Alignment",
			"Ca&tegory",
		},
		{ // 9
			"T&ype",
			"&Title",
			"&Forenames",
			"&Surname",
			"&Company",
			"&Email",
			"&Phone",
			"&Mobile",
			"U&RL",
			"F&ind",
			"&New",
			"&Duplicate",
			"&Update",
			"De&lete",
			"Confi&gure",
			"St&atistics",
			"E&xport",
			"&Open",
			"Sa&ve",
			"&Quit",
		},
		{ // 10
			"&Remaining Issues",
			"T&ype",
			"&Title",
			"&Forenames",
			"Surname",
			"&Company",
			"&Email",
			"&Phone",
			"&Mobile",
			"&URL",
			"F&ind",
			"&New",
			"&Duplicate",
			"Upd&ate",
			"De&lete",
			"Confi&gure",
			"&Statistics",
			"E&xport",
			"&Open",
			"Sa&ve",
			"&Quit",
		},
		{ // 11
			"Initial &No",
			"&Max Supportable No",
			"&Initial Memory Size",
			"Mutation &Rate",
			"&Survival Rate",
			"S&tart",
			"Pa&use",
			"&Zoom",
		},
		{ // 12
			"&General",
			"&Remember Geometry",
			"Start With Editor &Window",
			"Start With &Shell Window",
			"Automati&cally Hide the Find and Replace dialog",
			"&Input prompt",
			"&Output prompt",
			"&Python executable",
			"Working &directory",
			"Show Window Information at Start&up",
			"&Max Tooltip Length",
			"Max &Lines to Scan",
			"&Highlighting",
			"&Tab width",
			"&Font",
			"On &New File",
			"&At Startup",
			"P&ython Documentation",
		},
		{ // 13
			"Undo",
			"&Redo",
			"&Copy",
			"Cu&t",
			"&Paste",
			"Select &All",
			"Comp&lete",
			"&Find",
			"Find &Next",
			"R&eplace Next",
			"&Goto Line",
			"Goto &Matching",
			"&Indent Region",
			"Unin&dent Region",
			"C&omment Region",
			"&Uncomment Region",
		},
		{ // 14
			"Add &Top-Level Term",
			"Add Sub-Term &of",
			"Add &Child-Term of",
			"&Add Sibling-Term of",
			"&Find",
			"&Insert Unicode Character",
			"&Bold",
			"Ita&lic",
			"&Underline",
			"Stri&ke-out",
			"Su&perscript",
			"&Subscript",
			"&No Superscript or Subscript",
			"&Monospace",
			"Sans S&erif",
			"Se&rif",
		},
		{ // 15
			"&Toolbars",
			"Status &Bar",
			"S&idebar",
			"&Stop",
			"&Reload",
			"T&ext Size",
			"&Page Style",
			"&Character Encoding",
			"P&age Source",
			"&Full Screen",
		},
		{ // 16
			"&New",
			"&Open...",
			"Recent Doc&uments",
			"&Wizards",
			"&Close",
			"Save",
			"Save &As...",
			"Save A&ll",
			"&Reload",
			"&Versions...",
			"&Export...",
			"Export as PD&F...",
			"&Send",
			"Properties...",
			"&Digital Signatures...",
			"&Templates",
			"Preview in Web &Browser",
			"Pa&ge Preview",
			"Pr&int...",
			"&Printer Settings...",
			"E&xit",
		},
		{ // 17
			"&Undo",
			"&Redo",
			"Cut",
			"Cop&y",
			"Paste",
			"&Paste Special...",
			"Select &Text",
			"&Select All",
			"&Changes",
			"Co&mpare Documents...",
			"Find and Replace",
			"&Navigator",
			"&AutoText...",
			"&Exchange Database...",
			"Fiel&ds...",
			"&Footnote...",
			"Inde&x Entry...",
			"&Bibliography Entry",
			"&Hyperlink",
			"&Links...",
			"Plu&g-in",
			"&ImageMap",
			"&Object",
		},
		{ // 18
			"&Print Layout",
			"&Web Layout",
			"&Toolbars",
			"&Status Bar",
			"&Input Method Status",
			"&Ruler",
			"Text &Boundaries",
			"&Field Shadings",
			"Fi&eld Names",
			"&Nonprinting Characters",
			"&Hidden Paragraphs",
			"&Data Sources",
			"F&ull Screen",
			"&Zoom",
		},
		{ // 19
			"&Manual Break",
			"Fiel&ds",
			"Special Character...",
			"Formatting Mar&k",
			"Section...",
			"H&yperlink",
			"&Header",
			"Footer",
			"&Footnote",
			"C&aption",
			"&Bookmark",
			"&Cross-reference...",
			"&Note...",
			"&Script...",
			"&Indexes and Tables",
			"&Envelope...",
			"F&rame...",
			"&Table...",
			"Hori&zontal Ruler...",
			"&Picture",
			"Mo&vie and Sound",
			"&Object",
			"Floatin&g Frame",
			"Fi&le...",
		},
		{ // 20
			"&Default Formatting",
			"&Character...",
			"&Paragraph...",
			"&Bullets and Numbering...",
			"Pag&e...",
			"&Title Page",
			"C&hange Case",
			"Co&lumns",
			"&Sections...",
			"St&yles and Formatting",
			"A&utoFormat",
			"A&nchor",
			"&Wrap",
			"&Alignment",
			"A&rrange",
			"&Flip",
			"&Group",
			"&Object",
			"Fra&me...",
			"P&icture...",
		},
		{ // 21
			"&Insert",
			"&Delete",
			"S&elect",
			"&Merge Cells",
			"&Split Cells",
			"&Protect Cells",
			"Me&rge Table",
			"Sp&lit Table",
			"&AutoFormat...",
			"A&utofit",
			"&Heading rows repeat",
			"&Convert",
			"S&ort...",
			"&Formula",
			"&Number Format...",
			"Table &Boundaries",
			"&Table Properties...",
		},
		{ // 22
			"&Spellcheck...",
			"&Language",
			"&Word Count",
			"&AutoCorrect...",
			"Outline &Numbering...",
			"&Footnotes...",
			"&Gallery",
			"Media &Player",
			"&Bibliography Database",
			"Ma&il Merge Wizard...",
			"So&rt...",
			"&Calculate",
			"&Update",
			"&Macros",
			"&Extension Manager...",
			"&XML Filter Settings...",
			"Cus&tomize...",
			"&Options...",
		},
		{ // 23
			"&Undo",
			"&Redo",
			"Cu&t",
			"&Copy",
			"&Paste",
			"Paste Special...",
			"Select Te&xt",
			"&Select All",
			"Changes",
			"Co&mpare Documents...",
			"Find and Replace",
			"&Navigator",
			"&AutoText...",
			"&Exchange Database...",
			"Fiel&ds...",
			"&Footnote...",
			"Index Entr&y...",
			"&Bibliography Entry",
			"&Hyperlink",
			"&Links...",
			"Plu&g-in",
			"&ImageMap",
			"&Object",
		},
		{ // 24
			"&Fields",
			"Specia&l Character...",
			"Formatting Mar&k",
			"Section...",
			"H&yperlink",
			"Hea&der",
			"Footer",
			"Footnote",
			"C&aption",
			"&Bookmark",
			"&Cross-reference...",
			"&Note...",
			"&Script...",
			"&Indexes and Tables",
			"&Envelope...",
			"F&rame...",
			"&Table...",
			"&Horizontal Ruler...",
			"&Picture",
			"&Movie and Sound",
			"&Object",
			"Floatin&g Frame",
		},
		{ // 25
			"&Undo",
			"Redo",
			"&Repeat",
			"Cu&t",
			"&Copy",
			"Paste",
			"Paste Special...",
			"Select All",
			"Chan&ges",
			"Co&mpare Document...",
			"Find &and Replace...",
			"&Navigator",
			"&Headers and Footers...",
			"&Fill",
			"D&elete Contents...",
			"&Delete Cells...",
			"&Sheet",
			"Delete Manual &Break",
			"&Links...",
			"&Plug-in",
			"&ImageMap",
			"&Object",
		},
		{ // 26
			"&Undo",
			"Redo",
			"&Repeat",
			"Cu&t",
			"&Copy",
			"Paste",
			"Paste Special...",
			"Select All",
			"Chan&ges",
			"Co&mpare Document...",
			"Find &and Replace...",
			"&Navigator",
			"&Headers and Footers...",
			"&Fill",
			"D&elete Contents...",
			"&Delete Cells...",
			"&Sheet",
			"Delete Manual &Break",
			"&Links...",
			"&Plug-in",
			"&ImageMap",
			"&Object",
		},
		{ // 27
			"Undo",
			"&Redo",
			"&Copy",
			"Cu&t",
			"&Paste",
			"Select &All",
			"Comp&lete",
			"&Find",
			"Find &Next",
			"R&eplace Next",
			"&Goto Line",
			"Goto &Matching",
			"&Indent Region",
			"Unin&dent Region",
			"C&omment Region",
			"&Uncomment Region",
		},
		{ // 28
			"&Undo",
			"&Redo",
			"&Copy",
			"Cu&t",
			"&Paste",
			"&Find",
			"Find &Again",
			"F&ind && Replace",
		},
		{"O&ne", "&Two"}, // 29
		{ // 30
			"&Undo",
			"&Redo",
			"&Copy",
			"Cu&t",
			"&Paste",
			"&Find",
			"Find &Again"},
	}

	for i := 0; i < len(originals); i++ {
		original := originals[i]
		expected := expecteds[i]
		hinted, _, err := Hinted(original)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		sanityCheck(hinted, t)
		for j := 0; j < len(original); j++ {
			if hinted[j] != expected[j] {
				t.Errorf("#%d", i)
				t.Errorf("expected %q, got %q", expected[j], hinted[j])
			}
		}
	}
}

func TestBad1(t *testing.T) {
	original := []string{
		"Undo",
		"Redo",
		"&Copy",
		"&Cut", // duplicate preset
		"Paste",
		"Find",
		"Find Again",
		"Find && Replace",
	}
	_, _, err := Hinted(original)
	if err == nil {
		t.Error("expected an error")
	}
	// TODO check the error
}

func TestBad2(t *testing.T) {
	original := []string{
		"Undo",
		"Redo",
		"Copy",
		"Cu&t",
		"Paste",
		"Find",
		"Find Again",
		"Find && Replace",
		"Undo",
		"Redo",
		"Copy",
		"Cut",
		"Paste",
		"Find",
		"Find Again",
		"Find && Replace",
		"Undo",
		"Redo",
		"Copy",
		"Cut",
		"Paste",
		"Find",
		"Find Again",
		"Find && Replace",
		"Undo",
		"Redo",
		"Copy",
		"Cut",
		"Paste",
		"Find",
		"Find Again",
		"Find && Replace",
		"Undo",
		"Redo",
		"Copy",
		"Cut",
		"Paste",
		"Find",
		"Find Again",
		"Find && Replace",
	}
	_, _, err := Hinted(original)
	if err == nil {
		t.Error("expected an error")
	}
	// TODO check the error
}

func sanityCheck(hinted []string, t *testing.T) {
	used := make(map[rune]bool, len(hinted))
	for _, hints := range hinted {
		hints := []rune(strings.ReplaceAll(hints, "&&", "||"))
		i := slices.Index(hints, '&')
		if i > -1 && i+1 < len(hints) {
			c := hints[i+1]
			_, found := used[c]
			if found {
				t.Errorf("unexpected duplicate %c in %s", c,
					strings.Join(hinted, "| "))
			}
			used[c] = true
		}
	}
}
