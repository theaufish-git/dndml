from flask import Flask, send_file

app = Flask(__name__)

@app.route('/dndml/<kind>/<version>')
def dndml(kind, version):
    print(kind,version)

    return send_file("{kind}/{version}.yaml".format(kind=kind,version=version))

app.run(host='0.0.0.0', port=8100)