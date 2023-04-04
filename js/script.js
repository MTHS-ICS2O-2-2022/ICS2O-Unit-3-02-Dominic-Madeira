// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: March 2023
'use strict'
/**
 * This function calculates area and perimeter of rectangle.
 */
function myButtonClicked () {
  // input
  const length = parseFloat(document.getElementById('length').value)
  const width = parseFloat(document.getElementById('width').value)
  const height = parseFloat(document.getElementById('height').value)


  // process
  const volume = (length * width * height) / 3

  // output
  document.getElementById('volume').innerHTML = 'The Volume is: ' + volume + ' mm³'
}