export default function ResultCard({ result }) {
  if (!result) return null

  return (
    <div className="card" style={{ marginTop: "20px" }}>
      <strong>Short URL</strong>
      <p>{result}</p>
    </div>
  )
}
