import csv
import os
from datetime import datetime
from supabase import create_client, Client

from dotenv import load_dotenv

import os 
dir_path = os.path.dirname(os.path.realpath(__file__))

load_dotenv(os.path.join(dir_path,'..', '.env'))

# Initialize Supabase client
url: str = os.getenv("SUPABASE_URL")
key: str = os.environ.get("SUPABASE_PRIVATE_KEY")


supabase: Client = create_client(url, key)

print(supabase)

# Define the CSV file to read
file = "ad_click_data.csv"
bulk_import_data =[]
count=0
# # Open the CSV file and insert each row into Supabase
with open(file, 'r') as ad_data_file:
    reader = csv.DictReader(ad_data_file)
    for row in reader:
        print(f"Inserting row with timestamp {row['timestamp']} and user_id {row['user_id']}")
        print(row)
        # Insert the row into the bulk import
        bulk_import_data.append(row)
        count+1
        if count > 0:
            print("Row inserted successfully.")
        else:
            print("Failed to insert row.")


print(bulk_import_data)
supabase.table('ad_data').insert(bulk_import_data).execute()