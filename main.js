document.addEventListener('DOMContentLoaded', (event) => {
    //the event occurred
    let root = document.getElementById("root")

    console.log(root)

    let img = new Image();

    img.src = "/img/file-folder-n_l-dark.svg";

    root.appendChild(img);
})

