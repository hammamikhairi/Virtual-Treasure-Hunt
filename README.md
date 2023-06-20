# Virtual Treasure Hunt - The Most Challenging Hunt Ever

Welcome to the incredible Virtual Treasure Hunt! Brace yourself for the most mind-boggling and frustrating adventure you'll ever experience. We've combined the cutting-edge technologies of Next.js for the frontend and Go for the backend to bring you this ultimate test of patience and wit. Good luck, you'll need it!

## Levels and Solutions - Prepare to Be Amazed...or Not

Oh, the levels we have in store for you! Get ready to tear your hair out as you attempt to solve these mind-bending puzzles. Here they are, along with their ingenious solutions (or lack thereof):

| Level | Title | Solution |
|-------|-------|----------|
| 1     | The Frustratingly Obvious | <details><summary>Solution</summary>The flag is in the source code of the page &lt;F12&gt;</details> |
| 2     | A Tower of Characters | <details><summary>Solution</summary>The file is a picture encoded in base64. You should use a converter to get the image containing the flag.</details> |
| 3     | Images Are More Than What They Seem | <details><summary>Solution</summary>Change the extension of the image to .html and open it in the browser.</details> |
| 4     | Server Down? | <details><summary>Solution</summary>Play the Dino game long enough and the Flag will appear.</details> |
| 5     | The Fat Image | <details><summary>Solution</summary>Extract the image using winrar (or any alternative compression software) then concat the 2 halfs to get the full QR code.</details> |
| 6     | Decoding the Unthinkable | <details><summary>Solution</summary>Reverse the audio to get the flag.</details> |
| 7     | A small Break | <details><summary>Solution</summary>The Flag is hidden in my GitHub's profile README.md</details> |
| 8     | Forgotten Treasures | <details><summary>Solution</summary>Each of the previous pages has different titles. When we join all the titles, we get the full flag</details> |
| 9     | The Stolen Flag | <details><summary>Solution</summary>The IP address and password of a virtual machine will be found in one of the files of system_logs.zip. When connected to and executing the command `sl`, we will see a train containing the final flag</details> |

## Want to Experience the Frustration Locally?

If you're brave enough (or masochistic enough), you can clone the local branch of this repository and test your puzzle-solving skills on your own machine. It's a whole new level of torture!

Clone the local version of the trea sure hunt :

```bash
git clone -b local --single-branch https://github.com/hammamikhairi/Virtual-Treasure-Hunt
```

start the backend (make sure you have [go installed](https://go.dev/doc/install) )

```bash
cd backend
go mod tidy
go run .
```

start the frondend (npm/yarn required)

```bash
cd Website
npm i
npm run dev
```
