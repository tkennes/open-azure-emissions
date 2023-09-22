## 
read -r -d '' DATA << EOM
    { "meterCategory":"Azure App Service",
    "MeterSubCategory":"Azure App Service Premium v2 Plan",
    "Product":"Azure App Service Premium v2 Plan - P1 v2 - EU West",
    "Quantity":0.2,
    "UnitOfMeasure": "1",
    "Location": "EU West"
    }
EOM

curl -d $DATA -H "Content-Type: application/json" -X POST localhost:9005/footprint
