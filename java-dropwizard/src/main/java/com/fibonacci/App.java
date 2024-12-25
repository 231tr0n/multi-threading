package com.fibonacci;

import io.dropwizard.core.Application;
import io.dropwizard.core.Configuration;
import io.dropwizard.core.setup.Bootstrap;
import io.dropwizard.core.setup.Environment;

public class App extends Application<Configuration> {
  public static void main(String[] args) throws Exception {
    new App().run(args);
  }

  @Override
  public void initialize(Bootstrap<Configuration> bootstrap) {
    super.initialize(bootstrap);
  }

  @Override
  public void run(Configuration configuration, Environment environment) {
    environment.jersey().register(new ComputePlatform());
    environment.jersey().register(new ComputeVirtual());
    environment.healthChecks().register("application", new ApplicationHealthCheck());
  }
}
