package com.fibonacci;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class AppTest {

  @Test
  public void fibonacciTest() {
    assertEquals(0, App.compute(0));
    assertEquals(1, App.compute(1));
    assertEquals(2, App.compute(2));
    assertEquals(4, App.compute(3));
    assertEquals(7, App.compute(4));
    assertEquals(12, App.compute(5));
  }
}
