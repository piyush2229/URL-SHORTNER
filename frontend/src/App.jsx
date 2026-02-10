import { useState } from "react"
import Header from "./components/Header"
import ShortenerForm from "./components/ShortenerForm"
import ResultCard from "./components/ResultCard"
import Footer from "./components/Footer"
import "./styles/app.css"
import "./styles/responsive.css"

export default function App() {
  const [result, setResult] = useState("")

  return (
    <div className="container">
      <Header />
      <ShortenerForm onResult={setResult} />
      <ResultCard result={result} />
      <Footer />
    </div>
  )
}
