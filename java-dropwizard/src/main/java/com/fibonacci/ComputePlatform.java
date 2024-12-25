package com.fibonacci;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.QueryParam;
import jakarta.ws.rs.core.MediaType;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

@Path("/platform")
public class ComputePlatform {
  /**
   * Compute sum of fibonacci series till certain number.
   *
   * @param n Number till which fibonacci series sum is cacluated.
   * @return the total amount.
   */
  @GET
  @Produces(MediaType.TEXT_PLAIN)
  public static long compute(@QueryParam("n") int n) {
    final List<Worker> workers = new ArrayList<Worker>();

    for (int i = 0; i <= n; i++) {
      workers.add(new Worker(i));
    }

    ExecutorService threadPool = Executors.newCachedThreadPool(Thread.ofPlatform().factory());
    for (final Worker worker : workers) {
      threadPool.execute(worker);
    }
    threadPool.shutdown();
    while (!threadPool.isTerminated()) {}

    long total = 0;
    for (final Worker worker : workers) {
      total += worker.answer();
    }
    return total;
  }
}
