# TDT web app
This is a TDT project

## newsData.py
- Process newsData.json and store in db/gms/month monthly

## parser.py
- Get response of the entity extraction from DBPedia
```
python parser.py 1
```
> The second option is discarded since the DBPedia now changes the structure of the response

## plsa.py
- This is a pLsa procedure that generating the model

## web
- The web application can run separately once the model contains the default content(rubbish model), entity files(parser.py 1), news data in month(newsdata.py)
