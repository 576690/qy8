let tableobj = document.getElementById("index-table")

// 添加任务函数
function add() {
    let taskname = document.getElementById("taskname").value
    if (taskname == '') {
        alert("留言不能为空哦！")
    } else {
        let task = document.createElement("tr")
        let idNum = "tr" + Date.now()
        task.setAttribute("id", idNum);
        let childNode = '<td>' + taskname + '</td><td><button onclick="complete(' + idNum + ')">删除</button></td>';
        task.innerHTML = childNode;
        tableobj.appendChild(task);
        document.getElementById("taskname").value = ''
        localStorage.setItem(idNum, taskname)
    }
}

// 任务完成函数
function complete(idNum) {
    idNum.style.display = "none"
    localStorage.removeItem(idNum.id)
}

//页面初始化的时候，把本地数据生成列表
// function load() {
//     var arr = {}
//     for (let i = 0; i < localStorage.length; i++) {
//         var key = localStorage.key(i);
//         arr[key] = localStorage.getItem(key)
//     }

//     for (let i in arr) {
//         let taskname = localStorage.getItem(i);
//         let task = document.createElement("tr")
//         task.setAttribute("id", i);
//         let childNode = '<td>' + taskname + '</td><td><button onclick="complete(' + i + ')">完成</button></td>';
//         task.innerHTML = childNode;
//         tableobj.appendChild(task);
//     }
// }
load();