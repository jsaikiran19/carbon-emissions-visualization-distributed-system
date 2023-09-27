from flask import Flask, request, jsonify
import requests
from flask_cors import CORS
import os
import dotenv
import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import pandas
from producer import produce
from datetime import datetime
app = Flask(__name__)
CORS(app)
@app.route("/")
def start():
    return "App running"

def create_plot(X,y, view_by):
     if view_by=='stateId':
        fig, ax = plt.subplots()
        ax.bar(X, y)
        
        ax.set_xlabel(view_by)
        ax.set_ylabel('CO2 Emission value')
        # plt.show()
        # img = PIL.Image.frombytes('RGB', fig.canvas.get_width_height(),fig.canvas.tostring_rgb())
        # print(img)
        plt.savefig('co2.jpg', format='jpg')
        # bucket.blob('co2.png').upload_from_string(img, 'png')
        plt.close()

@app.route('/plot', methods=['POST'])
def plotData():
    body = (request.get_json())
    facets = body.get('facets',None)
    view_by = body.get('viewBy',None)
    headers = {
        'api_key': os.getenv('API_KEY'),
        'data[0]': "value",
        'start':'2015',
        'length':'1000',
    }
    for key in facets:
        for i,state in enumerate(facets[key]):
            headers["facets[{0}][{1}]".format(key,i)] = state
    os.environ['GOOGLE_APPLICATION_CREDENTIALS'] = os.getenv('GCLOUD_JSON')
    req = requests.get('https://api.eia.gov/v2/co2-emissions/co2-emissions-aggregates/data/?frequency=annual', headers)
    data = req.json()['response']['data']
    df = pandas.DataFrame(data)
    X = df[view_by].unique()
    y = df.groupby([view_by])['value'].sum()
    # client = storage.Client()
    # bucket = client.get_bucket(os.getenv('BUCKET_NAME'))
    create_plot(X,y,view_by)
    produce(str(datetime.now()))

    return jsonify({"message":"Plot created and placed in topic successfully","status":"Success"})

if __name__=='__main__':
    dotenv.load_dotenv()
    host = os.getenv('HOST')
    port = os.getenv('PORT')
    app.run(host, port)