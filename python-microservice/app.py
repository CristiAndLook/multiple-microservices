from flask import Flask

app = Flask(__name__)

@app.route('/get')
def get():
    return 'Everything is fine!'