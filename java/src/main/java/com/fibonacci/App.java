package com.fibonacci;

import com.sun.net.httpserver.HttpServer;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.net.URI;
import java.net.URISyntaxException;
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

  public static void main(String[] args) throws IOException {
    HttpServer server = HttpServer.create(new InetSocketAddress(8081), 0);
    server.createContext(
        "/",
        exchange -> {
          exchange.getResponseHeaders().add("Content-Type", "text/plain; charset=UTF-8");
          OutputStream out = exchange.getResponseBody();
          try {
            URI uri = new URI(exchange.getRequestURI().toString());
            String query = uri.getQuery();
            byte[] response =
                Long.toString(App.compute(Integer.parseInt(query.split("=")[1]))).getBytes("UTF-8");
            exchange.sendResponseHeaders(400, response.length);
            out.write(response);
          } catch (URISyntaxException | NumberFormatException e) {
            byte[] response = "Wrong parameter 'n'".getBytes("UTF-8");
            exchange.sendResponseHeaders(400, response.length);
            out.write(response);
          }
          out.close();
        });
    server.start();
  }
}
