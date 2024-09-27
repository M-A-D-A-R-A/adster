import csv
import random
import socket
import time
from datetime import datetime, timedelta
import faker

# Initialize Faker to generate random data
fake = faker.Faker()

# Constants for random data generation
num_records = 5
countries = ["US", "CA", "GB", "AU", "DE", "FR", "IN", "CN", "JP", "BR"]
regions = ["California", "Ontario", "London", "New South Wales", "Bavaria", "Île-de-France", "Maharashtra", "Beijing", "Tokyo", "São Paulo"]
cities = ["Los Angeles", "Toronto", "London", "Sydney", "Munich", "Paris", "Mumbai", "Shanghai", "Tokyo", "Rio de Janeiro"]
device_types = ["desktop", "mobile", "tablet"]
operating_systems = ["Windows", "macOS", "Linux", "Android", "iOS"]
browsers = ["Chrome", "Firefox", "Safari", "Edge", "Opera"]
domains = ["example.com", "testsite.com", "demo.com", "sample.org", "mywebsite.net"]
ad_positions = ["top", "middle", "bottom", "sidebar"]
ad_sizes = ["300x250", "728x90", "160x600", "300x600", "970x250"]

def generate_ip():
    return f"{random.randint(1, 255)}.{random.randint(0, 255)}.{random.randint(0, 255)}.{random.randint(0, 255)}"

def generate_timestamp(start_date, end_date):
    return start_date + timedelta(seconds=random.randint(0, int((end_date - start_date).total_seconds())))

# Set the date range for timestamps
start_date = datetime.now() - timedelta(days=30)
end_date = datetime.now()

# Open CSV file to write data
with open('ad_click_data.csv', 'w', newline='') as csvfile:
    csvwriter = csv.writer(csvfile)
    # Write header
    csvwriter.writerow(['timestamp', 'ip', 'user_id', 'geo_country', 'geo_region', 'geo_city', 
                        'device_type', 'os', 'browser', 'domain', 'url', 'ad_position', 'ad_size'])
    
    for _ in range(num_records):
        timestamp = generate_timestamp(start_date, end_date).isoformat()
        ip = generate_ip()
        user_id = fake.uuid4()
        geo_country = random.choice(countries)
        geo_region = random.choice(regions)
        geo_city = random.choice(cities)
        device_type = random.choice(device_types)
        os = random.choice(operating_systems)
        browser = random.choice(browsers)
        domain = random.choice(domains)
        url = f"https://{domain}/page/{random.randint(1, 100)}"
        ad_position = random.choice(ad_positions)
        ad_size = random.choice(ad_sizes)

        # Write the generated row to CSV
        csvwriter.writerow([timestamp, ip, user_id, geo_country, geo_region, geo_city, 
                            device_type, os, browser, domain, url, ad_position, ad_size])
        
        print(f"Generated record for user_id: {user_id}")

print("Data generation completed!")
