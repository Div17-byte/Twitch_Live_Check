First you need to get Twitch ClientId, Auth Token.
Step 1: Go to dev.twitch.tv/console/apps
Step 2: Login with your credentials
Step 3: Go to Application Tab and Register an application, 
    a> Use any valid name
    b> in OAuth Redirect URLs put this -> https://localhost:4200
    c> select any category

Step 4: After you register click on Manage you will see you Client ID and Client Secret 
Step 5: Click New Secret to generate new Secret 

Step 6: 
You will see REPLACE_ME_WITH_CLIENTID in below link, replace it with your clientId which you got it from twitch when you registered your application,
DO SAME WITH CLIENT SECRET !!
https://id.twitch.tv/oauth2/token?client_id=REPLACE_ME_WITH_CLIENT_ID&client_secret=REPLACE_ME_WITH_CLIENT_SECRET&grant_type=client_credentials

Step 7: After you replaced client_id and client_secret  Go to https://reqbin.com/

a> Paste your link in the BOX
b> select "POST" 

then click on Send

You will see the Response on the right

copy access_token value, which should look like this -> tyumkmktrr8jnx0bvmnhe480e2hsb3

EDIT Twitch_LiveChecker.go file on top you will see two variables named clientId and authToken  
cliendId = your client Id 
authToken = value of access token

Step 8: run "go build Twitch_LiveChecker.go" without the quotes in the terminal
NOTE* install go if you don't have it installed to use the go build command https://golang.org/doc/install

Step 9: After the build is finished you will the Twitch_LiveChecker.exe file in the same location


HOW TO USE
1. Edit "input.dat" file as text using any text editor
2. Paste Streamer name seprating each name with a comma ","
example.
shroud,bnans,Mizkif
3. Save file
4. Run the program
5. It will automatically open Live Stream if the streamer is Live on Twitch
