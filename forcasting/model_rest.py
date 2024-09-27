from flask import Flask, request, jsonify
import numpy as np
import xgboost as xgb
import pandas as pd
from clickhouse_driver import Client
import jwt
import os 

from functools import wraps
from typing import Callable
from dotenv import load_dotenv

# Initialize the Flask app
app = Flask(__name__)
dir_path = os.path.dirname(os.path.realpath(__file__))

load_dotenv(os.path.join(dir_path, '.env'))


key = os.getenv("SECRET_KEY", "adster")

def token_required() -> Callable:
    def static_jwt_wrapper(view_function):
        @wraps(view_function)
        def wrapper(*args, **kwargs):
            # Check if the user is logged in
            status = Auth.get_logged_in_user(request)
            if not status:
                return {"error": "Invalid token"}, 401
            
            # Call the original view function if the token is valid
            return view_function(*args, **kwargs)

        return wrapper

    return static_jwt_wrapper

class Auth:
    @staticmethod
    def decode_auth_token(auth_token):
        try:
            # Decode the token using the secret key
            payload = jwt.decode(auth_token, key, algorithms=["HS256"])
            return payload
        except jwt.ExpiredSignatureError:
            return {"error": "Signature expired. Please log in again."}
        except jwt.InvalidTokenError:
            return {"error": "Invalid token. Please log in again."}

    @staticmethod
    def get_logged_in_user(new_request):
        # Get the auth token from the request headers
        auth_token = new_request.headers.get("Authorization")
        if auth_token:
            # Remove 'Bearer ' prefix if present
            if auth_token.startswith("Bearer "):
                auth_token = auth_token.split(" ")[1]
            resp = Auth.decode_auth_token(auth_token)
            # Check for errors in decoding
            if isinstance(resp, dict) and "error" in resp:
                return False
            
            identity = resp.get("claims", {}).get("identity", resp.get("identity"))
            system_identity = os.getenv("SYSTEM_IDENTITY")

            # Validate the identity
            if identity != system_identity:
                return False
            return True
        else:
            return False


@app.route('/forecast', methods=['POST'])
@token_required()
def forecast():
    # Get data from request
    data = request.json
    # breakpoint()
    # Extract geo targeting information
    user_age =[]
    target_id =[]
    for d in data.get('geo_target', {}).get('included', []):
        user_age.append(d.get('age',None))

    for t in data.get('geo_target', {}).get('included', []):
        target_id.append(t.get('target_id',None))
    
    device = data.get('device_type', {}).get('included', [])



    # Calculate daily impressions and reach based on your model
    if target_id and device:
        # Example calculation for daily impressions and reach
        base_impressions = 1000000  # Base value
        base_reach = 500000  # Base value

        daily_impressions = base_impressions * len(device)
        daily_reach = base_reach * len(target_id) 
    else:
        daily_impressions = 0
        daily_reach = 0

    # # Create feature array for prediction
    features = np.array([[np.mean(user_age), np.mean(target_id), np.mean(device)]])

    # Convert to DMatrix
    dtest = xgb.DMatrix(features)
    # Make prediction
    y_pred_prob = model.predict(dtest)
    # y_pred = 1 if y_pred_prob[0] > 0.5 else 0
    # y_pred=1
    # Prepare forecast result based on the prediction
    result = {
        "forecast": {
            "daily_impressions": daily_impressions,
            "daily_reach": daily_reach,
            "prediction": 0.85
        }
    }

    print(result)

    return jsonify(result)



if __name__ == '__main__':
   
    # Load the trained XGBoost model
    model = xgb.Booster()
    model.load_model('xgboost_model.json')

    app.run(debug=True)
