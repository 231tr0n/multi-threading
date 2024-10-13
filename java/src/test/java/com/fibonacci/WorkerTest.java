package com.fibonacci;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class WorkerTest {

  @Test
  public void fibonacciTest() {
    assertEquals(0, Worker.fibonacci(0));
    assertEquals(1, Worker.fibonacci(1));
    assertEquals(1, Worker.fibonacci(2));
    assertEquals(2, Worker.fibonacci(3));
    assertEquals(3, Worker.fibonacci(4));
    assertEquals(5, Worker.fibonacci(5));
  }
}
