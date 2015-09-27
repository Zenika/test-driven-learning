package kata

import "testing"

func Test_should_print_string_representation_when_passing_an_integer(t *testing.T) {

	rt := Display(1)

	expected := "1"
	if rt != expected {
		t.Errorf("Must display %s but display %s", expected, rt)
	}
}

// Pour les multiples de trois afficher "Fizz".
func Test_should_print_Fizz_when_passing_3(t *testing.T) {

	rt := Display(3)

	expected := "Fizz"
	if rt != expected {
		t.Errorf("Must display %s but display %s", expected, rt)
	}
}

// Pour les multiples de trois afficher "Fizz".
func Test_should_print_Fizz_when_passing_6(t *testing.T) {

	rt := Display(6)

	expected := "Fizz"
	if rt != expected {
		t.Errorf("Must display %s but display %s", expected, rt)
	}
}

// Pour les multiples de cinq afficher "Buzz".
func Test_should_print_Buzz_when_passing_5(t *testing.T) {

	rt := Display(5)

	expected := "Buzz"
	if rt != expected {
		t.Errorf("Must display %s but display %s", expected, rt)
	}
}

// Pour les multiples de cinq afficher "Buzz".
func Test_should_print_Buzz_when_passing_20(t *testing.T) {

	rt := Display(20)

	expected := "Buzz"
	if rt != expected {
		t.Errorf("Must display %s but display %s", expected, rt)
	}
}

// Pour les multiples de trois et de cinq afficher "FizzBuzz".
func Test_should_print_FizzBuzz_when_passing_15(t *testing.T) {

	rt := Display(15)

	expected := "FizzBuzz"
	if rt != expected {
		t.Errorf("Must display %s but return is %s", expected, rt)
	}
}

// Pour les multiples de trois et de cinq afficher "FizzBuzz".
func Test_should_print_FizzBuzz_when_passing_30(t *testing.T) {

	rt := Display(30)

	expected := "FizzBuzz"
	if rt != expected {
		t.Errorf("Must display %s but return is %s", expected, rt)
	}
}

// Afficher les chiffres de 1 to 100.
func Test_should_display_a_valid_result_when_diplay_numbers_between_1_and_100(t *testing.T) {

	rt := FizzBuzz(100)

	expected := "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n16\n17\nFizz\n19\nBuzz\nFizz\n22\n23\nFizz\nBuzz\n26\nFizz\n28\n29\nFizzBuzz\n31\n32\nFizz\n34\nBuzz\nFizz\n37\n38\nFizz\nBuzz\n41\nFizz\n43\n44\nFizzBuzz\n46\n47\nFizz\n49\nBuzz\nFizz\n52\n53\nFizz\nBuzz\n56\nFizz\n58\n59\nFizzBuzz\n61\n62\nFizz\n64\nBuzz\nFizz\n67\n68\nFizz\nBuzz\n71\nFizz\n73\n74\nFizzBuzz\n76\n77\nFizz\n79\nBuzz\nFizz\n82\n83\nFizz\nBuzz\n86\nFizz\n88\n89\nFizzBuzz\n91\n92\nFizz\n94\nBuzz\nFizz\n97\n98\nFizz\n"

	if rt != expected {
		t.Errorf("Must display %s but return is %s", expected, rt)
	}
}
