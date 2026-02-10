import { useState } from "react"
import "../styles/form.css"

export default function ShortenerForm({ onResult }) {
  const [url, setUrl] = useState("")
  const [loading, setLoading] = useState(false)

 const submit = async () => {
  if (!url) {
    alert("Enter a URL")
    return
  }

  try {
    setLoading(true)
    const res = await fetch("http://localhost:3000/shorten", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ url }),
    })

    if (!res.ok) {
      const err = await res.text()
      alert(err)
      return
    }

    const data = await res.json()
    onResult?.(data.short_url)
  } finally {
    setLoading(false)
  }
}


  return (
    <div className="card">
      <input
        className="input"
        placeholder="Paste your long URL here"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />
      <button className="button" onClick={submit}>
        {loading ? "Shortening..." : "Shorten URL"}
      </button>
    </div>
  )
}
