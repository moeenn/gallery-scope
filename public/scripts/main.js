function select(selector) {
  const elements = document.querySelectorAll(selector)
  if (!elements) throw new Error(`Elements with selector ${selector} not found`)

  return elements
}

const API = {
  async request(url, method, payload) {
    const res = await fetch(url, {
      method,
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
      },
      body: JSON.stringify(payload)
    })

    if (!res.ok) {
      throw new Error("Request failed")
    }

    return await res.json()
  },

  async post(URL, form) {
    const data = await this.request(URL, "POST", form)
      .catch(err => console.error(err))

    return data
  }
}

const notification = {
  show(text) {
    const [n] = select("#notification")

    n.querySelector("span").innerHTML = text
    n.classList.remove("hidden")

    setTimeout(this.hide, 5000)
  },

  hide() {
    const [n] = select("#notification")
    n.classList.add("hidden")
  }
}

const form = {
  getForm() {
    const inputs = select("input")
    let result = {}

    for (const input of inputs) {
      const type = input.type
      result[input.name] = (type == 'number')
        ? parseInt(input.value)
        : input.value
    }

    return result
  },

  clear() {
    const [base] = select("input")
    base.value = ""

    thumbs.clear()
  },

  async search() {
    const form = this.getForm()
    const data = await API.post("/api/urls", form)

    thumbs.clear()

    const selective = [
      data[0],
      data[data.length -1],
    ]

    for (const img of selective) {
      const thumb = thumbs.createThumb(img)
      thumbs.insertThumbnail(thumb)
    }
  },

  async download() {
    const form = this.getForm()
    await API.post("/api/download", form)
    notification.show("Starting downloads")
  },
}

const thumbs = {
  clear() {
    const [gallery] = select(".gallery")
    gallery.innerHTML = ""
  },

  createThumb({ full, thumb }) {
    const [template] = select("template#thumbnail")
    const clone = template.content.cloneNode(true)

    clone.querySelector("a").href = full
    clone.querySelector("img").src = thumb

    return clone
  },

  insertThumbnail(thumb) {
    const [target] = select(".container.gallery")
    target.appendChild(thumb)
  },
}

const app = { form, thumbs }
