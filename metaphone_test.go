package phonetics

import "testing"

func TestMetaphoneEmptyString(t *testing.T) {
	if EncodeMetaphone("") != "" {
		t.Errorf("Encode with empty string should return empty string")
	}
}
func TestMetaphoneLocalString(t *testing.T) {
	assertMetaphoneLocalEquals(t, "Donald", "TNLT", "en")
	assertMetaphoneLocalEquals(t, "Нижинский", "НИЖИНСК7", "ru")
	assertMetaphoneLocalEquals(t, "Нежинский", "НИЖИНСК7", "ru")
	assertMetaphoneLocalEquals(t, "Нежинський", "НИЖИНСК7", "ru")
	assertMetaphoneLocalEquals(t, "Нежинськый", "НИЖИНСК7", "ru")
	assertMetaphoneLocalEquals(t, "Неженськый", "НИЖИНСК7", "ru")
	assertMetaphoneLocalEquals(t, "Неженскый", "НИЖИНСК7", "ru")
	assertMetaphoneLocalEquals(t, "Ніженскый", "НИЖИНСК7", "ru")
}

func TestMetaphoneEncode(t *testing.T) {
	assertMetaphoneEquals(t, "Donald", "TNLT")
	assertMetaphoneEquals(t, "Zach", "SX")
	assertMetaphoneEquals(t, "Campbel", "KMPBL")
	assertMetaphoneEquals(t, "Cammmppppbbbeeelll", "KMPBL")
	assertMetaphoneEquals(t, "David", "TFT")
	assertMetaphoneEquals(t, "Wat", "WT")
	assertMetaphoneEquals(t, "What", "WT")
	assertMetaphoneEquals(t, "Gaspar", "KSPR")
	assertMetaphoneEquals(t, "ggaspar", "KSPR")

	assertMetaphoneEquals(t, "Нижинский", "НИЖИНСК7")
	assertMetaphoneEquals(t, "Нежинский", "НИЖИНСК7")
	assertMetaphoneEquals(t, "Нежинський", "НИЖИНСК7")
	assertMetaphoneEquals(t, "Нежинськый", "НИЖИНСК7")
	assertMetaphoneEquals(t, "Неженський", "НИЖИНСК7")
	assertMetaphoneEquals(t, "Неженскый", "НИЖИНСК7")
	assertMetaphoneEquals(t, "Ніженский", "НИЖИНСК7")
}

func assertMetaphoneEquals(t *testing.T, source string, target string) {
	if EncodeMetaphone(source) != target {
		t.Errorf("source doesn't match target. Input: %s, Result: %s, Target: %s", source, EncodeMetaphone(source), target)
	}
}
func assertMetaphoneLocalEquals(t *testing.T, source, target, local string) {
	if EncodeMetaphone(source, local) != target {
		t.Errorf("source doesn't match target. Input: %s, Result: %s, Target: %s", source, EncodeMetaphone(source, local), target)
	}
}
func TestMetaphoneWordsString(t *testing.T) {
	if words := EncodeMetaphoneWords("Нежинский Неженскый", "ru"); words[0] != "НИЖИНСК7" || words[1] != "НИЖИНСК7" {
		t.Errorf("source doesn't match target. Input: %s, Result: %s, Target: %s", "Нежинский Неженскый", EncodeMetaphoneWords("Нежинский Неженскый", "ru"), []string{"НИЖИНСК7", "НИЖИНСК7"})
	}
}
