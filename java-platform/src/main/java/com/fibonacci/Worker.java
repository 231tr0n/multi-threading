package com.fibonacci;

class Worker implements Runnable {
  private int question;
  private long ans;

  public Worker(int question) {
    this.question = question;
  }

  public long answer() {
    return ans;
  }

  public static long fibonacci(long n) {
    if (n == 0) {
      return 0;
    }
    if (n == 1) {
      return 1;
    }
    return Worker.fibonacci(n - 1) + Worker.fibonacci(n - 2);
  }

  @Override
  public void run() {
    this.ans = Worker.fibonacci(this.question);
  }
}
