package com.fibonacci;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class ComputeVirtualTest {

  @Test
  public void fibonacciTest() {
    assertEquals(0, ComputeVirtual.compute(0));
    assertEquals(1, ComputeVirtual.compute(1));
    assertEquals(2, ComputeVirtual.compute(2));
    assertEquals(4, ComputeVirtual.compute(3));
    assertEquals(7, ComputeVirtual.compute(4));
    assertEquals(12, ComputeVirtual.compute(5));
  }
}
