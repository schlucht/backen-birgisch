document.addEventListener('DOMContentLoaded', () => {
    const todos = document.querySelectorAll('p.todo')
    for (let td of todos) {
        td.addEventListener('click', (e) => {
            console.log(e.target.id)
        })
    }
})