package tamagotchi

import "testing"

func TestTamagotchi(t *testing.T) {
    expected := "Hello Tamagotchi!"
    actual := hi()
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
    }
}
