#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>
#include <DNSServer.h>
#include <ESP8266WebServer.h>
#include <WiFiManager.h>

// Set web server port number to 80
WiFiServer server(80);

// Variable to store the HTTP request

WiFiManager wifiManager;

const char* serverName = "http://103.67.186.22:8080/flow";

HTTPClient http;

//interval of post request
unsigned long interval = 5000;
volatile long timeState = 0;

//flow sensor data
int flowPin = 4; //GPIO05 D2
volatile long count;

//function called on interupt
void ICACHE_RAM_ATTR Flow(){
  count++;
}

void send_post(){
  if(WiFi.status()== WL_CONNECTED){
    WiFiClient client = server.available();
    String val = String(count);
    count = 0;   
    http.begin(client, serverName);
    String body = String("{\"value\":"+val+",\"device\":\"flow1\"}");
    http.addHeader("Content-Type", "application/json");
    int httpResponseCode = http.POST(body);
    Serial.print("HTTP Response code: ");
    Serial.println(httpResponseCode);
    http.end();
  }
}
void setup_wifi() {
  // connect wifi with blink indicator
  //192.168.4.1:80
  digitalWrite(BUILTIN_LED, HIGH);
  wifiManager.autoConnect("NodeMCU-Flow1");
}


void setup() {
  Serial.begin(9600);
  pinMode(BUILTIN_LED, OUTPUT);     // Initialize the BUILTIN_LED pin as an output
  pinMode(flowPin, INPUT_PULLUP);   //set D1 (GPIO05) as input  
  setup_wifi();
  attachInterrupt(digitalPinToInterrupt(flowPin), Flow, RISING);     // set interupt on flowpin 
}

void loop() {
  if(millis()-timeState > interval){
    if(WiFi.status()== WL_CONNECTED){
      digitalWrite(BUILTIN_LED, LOW);
      send_post(); //send post request
      timeState = millis();
    }
  }
}