package com.fibonacci;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class ComputePlatformTest {

  @Test
  public void fibonacciTest() {
    assertEquals(0, ComputePlatform.compute(0));
    assertEquals(1, ComputePlatform.compute(1));
    assertEquals(2, ComputePlatform.compute(2));
    assertEquals(4, ComputePlatform.compute(3));
    assertEquals(7, ComputePlatform.compute(4));
    assertEquals(12, ComputePlatform.compute(5));
  }
}
