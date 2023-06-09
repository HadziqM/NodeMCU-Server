#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>

const char* ssid = "CMCC-Kelapa 3";
const char* password = "gakngerti";
const char* serverName = "http://192.168.1.7:8080/flow";

bool connTest = true;

WiFiClient client;
HTTPClient http;

//interval of post request
unsigned long interval = 5000;
volatile long timeState = 0;

//flow sensor data
int flowPin = 5; //GPIO05 D1
volatile long count;

//function called on interupt
void ICACHE_RAM_ATTR Flow(){
  count++;
}

void send_post(double value){
  if(WiFi.status()== WL_CONNECTED){
    count = 0;   
    http.begin(client, serverName);
    String val = String(value,2);
    String body = String("{\"value\":\""+val+"\",\"device\":\"flow1\"}");
    http.addHeader("Content-Type", "application/json");
    int httpResponseCode = http.POST(body);
    Serial.print("HTTP Response code: ");
    Serial.println(httpResponseCode);
    http.end();
  }
}
void setup_wifi() {
  // connect wifi with blink indicator
  digitalWrite(BUILTIN_LED, LOW);
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  unsigned long status = 0; 
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    if (status % 2 == 0){
      digitalWrite(BUILTIN_LED, HIGH);
    }else{
      digitalWrite(BUILTIN_LED, LOW);
    }
    Serial.print(".");
  }
  randomSeed(micros());
  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  digitalWrite(BUILTIN_LED, HIGH);
  Serial.println(WiFi.localIP());
}


void setup() {
  Serial.begin(9600);
  pinMode(BUILTIN_LED, OUTPUT);     // Initialize the BUILTIN_LED pin as an output
  pinMode(flowPin, INPUT_PULLUP);   //set D1 (GPIO05) as input
  Serial.println("Ready");
  if (connTest){
    setup_wifi();
  }
  attachInterrupt(digitalPinToInterrupt(flowPin), Flow, RISING);     // set interupt on flowpin 
}

void loop() {
  if(millis()-timeState > interval){
    double value = (count * 1.68); //sensor spec 1.68mL per pulse
    if (connTest){
        send_post(value); //send post request
    }else{
        count = 0;
        Serial.print("Flow Counted in interval (mL): ");
        Serial.println(value);
    }
    timeState = millis();
  }
}