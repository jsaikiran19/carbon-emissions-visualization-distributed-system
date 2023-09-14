from flask import Flask
import os
import dotenv


app = Flask(__name__)

@app.route("/")
def start():
    return "App running"

if __name__=='__main__':
    dotenv.load_dotenv()
    host = os.getenv('HOST')
    port = os.getenv('PORT')
    app.run(host, port)