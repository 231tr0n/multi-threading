package com.fibonacci;

import java.util.ArrayList;
import java.util.List;

class App {
  public static long compute(int n) {
    final List<Worker> workers = new ArrayList<Worker>();

    for (int i = 0; i <= n; i++) {
      workers.add(new Worker(i));
    }

    final List<Thread> threads = new ArrayList<Thread>();
    for (final Worker worker : workers) {
      threads.add(new Thread(worker));
    }

    for (final Thread thread : threads) {
      thread.start();
    }

    for (final Thread thread : threads) {
      try {
        thread.join();
      } catch (InterruptedException e) {
        e.printStackTrace();
      }
    }

    long total = 0;
    for (final Worker worker : workers) {
      total += worker.answer();
    }
    return total;
  }

  public static void main(String[] args) {
    System.out.println("Total: " + App.compute(Integer.parseInt(args[1])));
  }
}
