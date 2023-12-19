# Messenger Clone V2

![](https://img.shields.io/badge/made%20by-DarienMiller-blue)
![](https://img.shields.io/badge/Golang-1.17-yellow)
![](https://img.shields.io/badge/React-red)
![](https://img.shields.io/badge/MongoDB-Cloud-green)
[![Netlify Status](https://api.netlify.com/api/v1/badges/1715535d-5c1e-46d4-80e2-fba5816cf2ca/deploy-status)](https://app.netlify.com/sites/messengerv2/deploys)

This is a Svelte implementation of my previous messenger clone I made this year. I was not satisfied with the code, so I'm starting fresh here! 

Features of this chat application include:
- A Public chat that all users can type in.
- Anonymous user sign in.
- Google + Github sign in (Later on).
- Private chats + user to user messaging.
- Image upload.
- Dark mode + Light mode toggle.

### Built With

* [Svelte](https://reactjs.org)
* [Fiber](https://github.com/gofiber/fiber)
* [Vite](https://vitejs.dev/)
* [PostgreSQL](https://www.postgresql.org/)
* [Netlify](https://bit.ly/3q4pcJz)
* [fly.io](https://fly.io/)
* [Pusher](https://pusher.com/)
* [Cloundinary](https://cloudinary.com/)

## Sign in form.
<img width="960" alt="signinform" src="https://github.com/darienmiller88/Better-Bank-Account/assets/32966645/6201e976-30fc-4b4c-90b5-9ea123036acd">

## Home
<img width="960" alt="home" src="https://github.com/darienmiller88/Better-Bank-Account/assets/32966645/fe390a7d-9c0c-4852-8367-86068d788d1c">

## Profile Form
<img width="960" alt="profileform" src="https://github.com/darienmiller88/Better-Bank-Account/assets/32966645/3258543b-ee28-4c86-80f6-89fd0acb8cb0">

## New Chat Form
<img width="960" alt="addnewchatform" src="https://github.com/darienmiller88/Better-Bank-Account/assets/32966645/aa0b261c-c7ba-4c94-9947-ae34b5b60b9e">

## Delete Message Form
<img width="960" alt="deletemessageform" src="https://github.com/darienmiller88/Better-Bank-Account/assets/32966645/36e993f0-9cc4-4693-9a08-8d00609c6f99">


### Requirements
* Clone the repository using `git clone https://github.com/darienmiller88/MessengerV2`
* Migrate the necessary information to your local `.env` as described in the `.env_sample` file
* Run `go build` to create a root level `MessengerV2.exe` file, and then run `.\MessengerV2` to run the executable. If an executable is not needed, simply input `go run main.go` instead, or `.\fresh` to enable a server restart on change.
* `cd` into the `client` folder, and run `npm run dev` to the Vite server, which should be on port 5173

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

<p align="right">(<a href="#top">back to top</a>)</p>

## License
[MIT](https://choosealicense.com/licenses/mit/)