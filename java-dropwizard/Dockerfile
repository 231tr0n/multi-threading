FROM eclipse-temurin:21-alpine

WORKDIR /root/trial

COPY target/fibonacci-1.0-SNAPSHOT.jar .

CMD ["java", "-Ddw.server.applicationConnectors[0].port=8083", "-jar", "fibonacci-1.0-SNAPSHOT.jar", "server"]
