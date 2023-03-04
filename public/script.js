function el(key) {
    return document.querySelector(key)
}

async function get(url, body) {
    const res = await fetch(url, { method: "POST", headers: { "Content-Type": "application/json" }, body: JSON.stringify(body) })

    return await res.json()
}


el("#send").addEventListener('submit', async (e) => {
    e.preventDefault()

    el("#response").textContent = ""

    const url = e.target.url.value

    const summarize = e.target.summarize.checked

    el("#loading").classList.add("show")

    const res = await get(`/ai/${summarize ? "summarize" : "explain"}`, {
        url: url,
        language: "en"
    })

    el("#loading").classList.remove("show")
    el("#response").textContent = res.result
})