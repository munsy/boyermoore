package boyermoore

import "testing"

func Test_BM_Struct(t *testing.T) {
	bm := NewBoyerMoore("WE HOLD THESE TRUTHS TO BE SELF-EVIDENT")

	// These should be found
	if !bm.Search("WE") {
		t.Errorf("Did not find \"WE\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("HOLD") {
		t.Errorf("Did not find \"HOLD\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("THESE") {
		t.Errorf("Did not find \"THESE\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("TRUTHS") {
		t.Errorf("Did not find \"TRUTHS\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("TO") {
		t.Errorf("Did not find \"TO\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("BE") {
		t.Errorf("Did not find \"BE\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("SELF") {
		t.Errorf("Did not find \"SELF\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("EVIDENT") {
		t.Errorf("Did not find \"EVIDENT\" in \"" + bm.ToSearch + "\"")
	}

	// These should not be found
	if bm.Search("aeiou") {
		t.Errorf("Found \"aeiou\" in \"" + bm.ToSearch + "\"")
	}
	if bm.Search("AEIOU") {
		t.Errorf("Found \"AEIOU\" in \"" + bm.ToSearch + "\"")
	}
	if bm.Search("OD") {
		t.Errorf("Found \"OD\" in \"" + bm.ToSearch + "\"")
	}
	if bm.Search("OH") {
		t.Errorf("Found \"OH\" in \"" + bm.ToSearch + "\"")
	}
	if bm.Search("  ") {
		t.Errorf("Found \"  \" in \"" + bm.ToSearch + "\"")
	}
	if bm.Search("") {
		t.Errorf("Found \"\" in \"" + bm.ToSearch + "\"")
	}

	// Special cases
	bm.ToSearch = ""
	if !bm.Search("") {
		t.Errorf("Didn't find \"\" in \"" + bm.ToSearch + "\"")
	}
	if bm.Search("A") {
		t.Errorf("Found \"A\" in \"" + bm.ToSearch + "\"")
	}
	bm.ToSearch = "A"
	if !bm.Search("A") {
		t.Errorf("Didn't find \"A\" in \"" + bm.ToSearch + "\"")
	}
	if bm.Search("") {
		t.Errorf("Found \"\" in \"" + bm.ToSearch + "\"")
	}

	// Other tests
	bm.ToSearch = "WE HOLD THESE TRUTHS TO BE SELF-EVIDENT"
	if !bm.Search("-") {
		t.Errorf("Did not find \"-\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("SELF-EVIDENT") {
		t.Errorf("Did not find \"SELF-EVIDENT\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("WE HOLD THESE TRUTHS TO BE SELF-EVIDENT") {
		t.Errorf("Did not find \"WE HOLD THESE TRUTHS TO BE SELF-EVIDENT\" in \"" + bm.ToSearch + "\"")
	}
	if !bm.Search("TRUTH") {
		t.Errorf("Did not find \"TRUTH\" in \"" + bm.ToSearch + "\"")
	}
}

func Test_bm(t *testing.T) {
	TestSearchString := "WE HOLD THESE TRUTHS TO BE SELF-EVIDENT"

	// These should be found
	if !bm(TestSearchString, "WE") {
		t.Errorf("Did not find \"WE\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "HOLD") {
		t.Errorf("Did not find \"HOLD\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "THESE") {
		t.Errorf("Did not find \"THESE\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "TRUTHS") {
		t.Errorf("Did not find \"TRUTHS\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "TO") {
		t.Errorf("Did not find \"TO\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "BE") {
		t.Errorf("Did not find \"BE\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "SELF") {
		t.Errorf("Did not find \"SELF\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "EVIDENT") {
		t.Errorf("Did not find \"EVIDENT\" in \"" + TestSearchString + "\"")
	}

	// These should not be found
	if bm(TestSearchString, "aeiou") {
		t.Errorf("Found \"aeiou\" in \"" + TestSearchString + "\"")
	}
	if bm(TestSearchString, "AEIOU") {
		t.Errorf("Found \"AEIOU\" in \"" + TestSearchString + "\"")
	}
	if bm(TestSearchString, "OD") {
		t.Errorf("Found \"OD\" in \"" + TestSearchString + "\"")
	}
	if bm(TestSearchString, "OH") {
		t.Errorf("Found \"OH\" in \"" + TestSearchString + "\"")
	}
	if bm(TestSearchString, "  ") {
		t.Errorf("Found \"  \" in \"" + TestSearchString + "\"")
	}
	if bm(TestSearchString, "") {
		t.Errorf("Found \"\" in \"" + TestSearchString + "\"")
	}

	// Special cases
	if !bm("", "") {
		t.Errorf("Didn't find \"\" in \"\"")
	}
	if !bm("A", "A") {
		t.Errorf("Didn't find \"A\" in \"A\"")
	}
	if bm("A", "") {
		t.Errorf("Found \"A\" in \"\"")
	}
	if bm("", "A") {
		t.Errorf("Found \"\" in \"A\"")
	}
	if !bm("", "") {
		t.Errorf("Didn't find \"\" in \"\"")
	}

	// Other tests
	if !bm(TestSearchString, "SELF-EVIDENT") {
		t.Errorf("Did not find \"SELF-EVIDENT\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "WE HOLD THESE TRUTHS TO BE SELF-EVIDENT") {
		t.Errorf("Did not find \"WE HOLD THESE TRUTHS TO BE SELF-EVIDENT\" in \"" + TestSearchString + "\"")
	}
	if !bm(TestSearchString, "TRUTH") {
		t.Errorf("Did not find \"TRUTH\" in \"" + TestSearchString + "\"")
	}
}
