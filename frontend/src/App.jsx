import { useState, useEffect } from 'react';
import MedicineForm from './MedicineForm';

function App() {
  const [medicines, setMedicines] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/api/medicines')
      .then((res) => {
        if (!res.ok) throw new Error('Gagal ambil data');
        return res.json();
      })
      .then((json) => {
        setMedicines(json.data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  const handleAddMedicine = (newMedicine) => {
    setMedicines((prev) => [...prev, newMedicine]);
  };

  return (
    <div className="min-h-screen bg-slate-50">
      <header className="bg-white border-b border-slate-200">
        <div className="max-w-5xl mx-auto px-6 py-4">
          <h1 className="text-xl font-bold text-slate-800">💊 Go Pharma</h1>
          <p className="text-sm text-slate-500">
            Sistem Manajemen Stok Obat
          </p>
        </div>
      </header>

      <main className="max-w-5xl mx-auto px-6 py-8">
        <MedicineForm onAddMedicine={handleAddMedicine} />

        <div className="bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden">
          <div className="px-6 py-4 border-b border-slate-200">
            <h2 className="text-lg font-semibold text-slate-800">
              Daftar Obat
            </h2>
          </div>

          {loading && (
            <p className="px-6 py-8 text-center text-slate-500 text-sm">
              Memuat data...
            </p>
          )}

          {error && (
            <p className="px-6 py-8 text-center text-red-500 text-sm">
              Error: {error}
            </p>
          )}

          {!loading && !error && (
            <table className="w-full text-sm">
              <thead>
                <tr className="bg-slate-50 text-left text-slate-500 text-xs uppercase tracking-wide">
                  <th className="px-6 py-3 font-medium">ID</th>
                  <th className="px-6 py-3 font-medium">Nama Obat</th>
                  <th className="px-6 py-3 font-medium">Kategori</th>
                  <th className="px-6 py-3 font-medium">Stok</th>
                  <th className="px-6 py-3 font-medium">Satuan</th>
                  <th className="px-6 py-3 font-medium">Harga</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-slate-100">
                {medicines.length === 0 && (
                  <tr>
                    <td
                      colSpan={6}
                      className="px-6 py-8 text-center text-slate-400"
                    >
                      Belum ada data obat
                    </td>
                  </tr>
                )}
                {medicines.map((m) => (
                  <tr key={m.id} className="hover:bg-slate-50">
                    <td className="px-6 py-3 text-slate-500">{m.id}</td>
                    <td className="px-6 py-3 text-slate-800 font-medium">
                      {m.name}
                    </td>
                    <td className="px-6 py-3 text-slate-600">
                      {m.category}
                    </td>
                    <td
                      className={`px-6 py-3 font-medium ${m.stock < 20 ? 'text-red-600' : 'text-slate-700'
                        }`}
                    >
                      {m.stock}
                    </td>
                    <td className="px-6 py-3 text-slate-600">{m.unit}</td>
                    <td className="px-6 py-3 text-slate-600">
                      Rp {m.price.toLocaleString('id-ID')}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          )}
        </div>
      </main>
    </div>
  );
}

export default App;