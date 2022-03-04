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

    const data = await res.json()
    return data
  },

  async post(URL, form) {
    const data = await this.request(URL, "POST", form)
      .catch(err => console.error(err))

    return data
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

  async search() {
    const form = this.getForm()
    const data = await API.post("/api/urls", form)

    thumbs.clear()

    for (const img of data) {
      const thumb = thumbs.createThumb(img)
      thumbs.insertThumbnail(thumb)
    }
  },

  async download() {
    const form = this.getForm()
    await API.post("/api/download", form)
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
