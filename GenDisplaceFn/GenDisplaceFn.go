/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/
package main

import (
	"fmt"
)

/*
Display function s = ½ a t^2 + vot + so,
returns a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement.
*/
func GenDisplaceFn(a_0, v_0, s_0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return ((0.5 * a_0 * (t * t)) + (v_0 * t) + s_0)
	}
	return fn
}

func main() {
	// input values
	var input_velocity, input_acceleration, input_init_displacement, input_time float64

	// Get input from user:
	fmt.Println("Enter Acceleration:")
	fmt.Scanln(&input_acceleration)

	fmt.Println("Enter velocity:")
	fmt.Scanln(&input_velocity)

	fmt.Println("Enter initial displacement:")
	fmt.Scanln(&input_init_displacement)

	fmt.Println("Enter value of time [s]:")
	fmt.Scanln(&input_time)

	// calculate distance
	displ1 := GenDisplaceFn(input_acceleration, input_velocity, input_init_displacement)
	// print out the results
	fmt.Println("Displacement after", input_time, "[s]")
	fmt.Println("for the given velocity of:", input_velocity, "[m/s]")
	fmt.Println("acceleration of:", input_acceleration, "[m/s^2]")
	fmt.Println("and initial displacement of:", input_init_displacement, "[m]")
	fmt.Println("equals to:", displ1(input_time), "[m]")

}
