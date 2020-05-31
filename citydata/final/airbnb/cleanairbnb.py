import pandas as pd
import copy

totaldata = pd.read_csv("airbnb_Manchester_long.csv")

def IsNumber(value):
    result = -10000
    try:
        result = float(value)
    except ValueError:
        pass
    return result

BedScore = {'Airbed': 0.8, 'Pull-out sofa': 0.6, 'Real Bed':1, 'Futon': 0.6}
BedRooms = {'Apartment': 1, 'Bed and breakfast': 1, 'Hotel room': 1, 'Private room': 1, 'Serviced apartment': 1}

def WashData(data):
    result = []
    for _, row in data.iterrows():
        oneline = {}
        beds = IsNumber(row['beds'])
        if beds == 1 or beds == 1.5:
            oneline['beds'] = beds
        else:
            continue

        if row['bed_type'] in BedScore:
            oneline['bed_type'] = row['bed_type']
            oneline['bed_score'] = BedScore[row['bed_type']]
        else:
            continue

        rating = IsNumber(row['review_scores_rating'])
        if rating >= 80:
            oneline['rating'] = rating
        else:
            continue
        
        if row['room_type'] not in BedRooms:
            continue
        price = row['price']
        if price.startswith('$'):
            price = price[1:]
        else:
            continue
        oneline['price'] = price
        
        latitude = IsNumber(row['latitude'])
        if latitude != -10000:
            oneline['latitude'] = latitude
        else:
            continue

        longitude = IsNumber(row['longitude'])
        if longitude != -10000:
            oneline['longtitude'] = longitude
        else:
            continue
        
        oneline['name'] = row['name']
        oneline['host_id'] = row['host_id']
        result.append(copy.deepcopy(oneline))
    
    return pd.DataFrame(result)

cleanData = WashData(totaldata)
cleanData.to_csv('airbnb_clean.csv')
print(cleanData.head())