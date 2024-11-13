import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class ServerPerSecondLoadSimulation extends Simulation {

  val httpProtocolEcho = http.baseUrl("http://echo-server:8081")
  val httpProtocolGin = http.baseUrl("http://gin-server:8081")

  val scnEcho = scenario("Echo Server Load Test")
    .exec(http("Echo Metrics").get("/ping"))
  val scnGin = scenario("Gin Server Load Test")
    .exec(http("Gin Metrics").get("/ping"))

  setUp(
    scnEcho.inject(
        constantUsersPerSec(2000).during(30)
    ).protocols(httpProtocolEcho),
    scnGin.inject(
        constantUsersPerSec(2000).during(30)
    ).protocols(httpProtocolGin)
  ).maxDuration(60.seconds)
}