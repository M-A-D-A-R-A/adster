import pandas as pd
from clickhouse_driver import Client
import xgboost as xgb
import numpy as np

# Connect to ClickHouse
client = Client('localhost')

def fetch_data():
    # select the data from the historic
    query = 'SELECT * FROM ad_training_data.training_data'

    data = client.query_dataframe(query)
    # breakpoint()
    return data

def train_model(data):
    # Assume data has columns ['age', 'ad_price', 'impressions', 'clicked']
    X = data[['age', 'target_id', 'device','daily_impressions','daily_reach']].values
    y = data['clicked'].values
    
    dtrain = xgb.DMatrix(X, label=y)
    param = {
        'max_depth': 3,
        'eta': 0.3,
        'objective': 'binary:logistic',
        'eval_metric': 'logloss',
    }
    num_round = 50
    bst = xgb.train(param, dtrain, num_round)
    return bst

def main():
    data = fetch_data()
    model = train_model(data)
    
    # Save the model (optional)
    model.save_model('xgboost_model.json')

if __name__ == '__main__':
    main()
