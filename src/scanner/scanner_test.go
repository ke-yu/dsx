package scanner

import(
    "testing"
)

func TestScanIdentifier(t *testing.T) {
    source := "foo"
    scanner := CreateScanner(source)

    token := scanner.Scan()

    expected := "foo"
    actual := token.StringValue()

    if (actual != expected) {
        t.Error("Failed at scanning identifier. Expected:", expected, ", got :", actual)
    }
}

func TestScanOctNumber(t *testing.T) {
    source := "077"
    scanner := CreateScanner(source)

    token := scanner.Scan()

    expected := "63"
    actual := token.StringValue()

    if (actual != expected) {
        t.Error("Failed at scanning identifier. Expected:", expected, ", got :", actual)
    }

    token = scanner.Scan()
}

func TestScanHexNumber(t *testing.T) {
    source := "0xff"
    scanner := CreateScanner(source)

    token := scanner.Scan()

    expected := "255"
    actual := token.StringValue()

    if (actual != expected) {
        t.Error("Failed at scanning identifier. Expected:", expected, ", got :", actual)
    }

    token = scanner.Scan()
}
