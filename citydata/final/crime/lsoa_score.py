import pandas as pd

monthdata = pd.read_csv("lsoa_crime_count_withmonth.csv")
yeardata = pd.read_csv("lsoa_crime_count_year.csv")
metalist = ["lsoa", "latitude", "longitude", "area", "month"]
crimes = ["Other theft", "Robbery", "Violence and sexual offences", "Burglary", "Public order", "Other crime", "Criminal damage and arson", "Vehicle crime", "Anti-social behaviour", "Shoplifting", "Bicycle theft", "Drugs", "Possession of weapons", "Theft from the person", "allcrime"]


scorelist = [{'Violence and sexual offences':1, 'Robbery': 1, 'Burglary': 1, "Bicycle theft":1, "Theft from the person":1}, 
            {'Public order':1, 'Criminal damage and arson': 1, 'Anti-social behaviour': 1, 'Shoplifting': 1}]

def CacuScore(data, scores):
    result = []
    count = 0
    for _, row in data.iterrows():
        if count % 1000 == 0:
            print("{}/{}".format(count, len(data)))
        count += 1
        oneline = {}
        for meta in metalist:
            oneline[meta] = row[meta]
        for index in range(len(scores)):
            scoreweight = scores[index]
            score = 0
            for key in scoreweight:
                score += row[key] * scoreweight[key]
            oneline["score_" + str(index)] = score
        result.append(oneline)
    return pd.DataFrame(result)
    
monthdf = CacuScore(monthdata, scorelist)
monthdf.to_csv("lsoa_score_month.csv")
yeardf = CacuScore(yeardata, scorelist)
yeardf.to_csv("lsoa_score_year.csv")