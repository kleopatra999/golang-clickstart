package main

import (
    "testing" //import go package for testing related functionality
    )



func Test_Add2Ints_2(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
   // t.Error("this is just hardcoded as an error.") //Indicate that this test failed and log the string as info
   t.Log("Everything is hunky-dory")
}