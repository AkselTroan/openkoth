from flask import Flask, session, url_for, redirect, render_template, request, abort, flash, send_from_directory, send_file
import requests

openkothAPI = "localhost:8080"
app = Flask(__name__)


@app.route("/")
def root():
    return render_template("index.html")


@app.route("/rooms")
def rooms():
    r = requests.get("http://" + openkothAPI + "/rooms")
    print(r.json()) # This method is convenient when the API returns JSON
    return r.json()


@app.route("/tmp")
def tmp():
    return render_template("rooms.html")



#@app.route("/NHIE", methods=["GET", "POST"])
#def never_have_i_ever():
#    if request.method == "POST":
#        lines = open('Resources/DrinkingGames/NeverHaveIEver/NeverHaveIEver.txt').read().splitlines()
#        statement = random.choice(lines)
#        return render_template("NHIE.html", statement=statement)
#    else:
#        return render_template("NHIE.html")



if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=9999)
