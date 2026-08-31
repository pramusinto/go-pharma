import { useState } from 'react';

function MedicineForm({ onAddMedicine }) {
    const [name, setName] = useState('');
    const [category, setCategory] = useState('');
    const [stock, setStock] = useState('');
    const [unit, setUnit] = useState('');
    const [price, setPrice] = useState('');
    const [submitting, setSubmitting] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!name || !category || !stock || !unit || !price) {
            alert('Semua field wajib diisi');
            return;
        }

        setSubmitting(true);

        try {
            const res = await fetch('http://localhost:8080/api/medicines', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name,
                    category,
                    stock: Number(stock),
                    unit,
                    price: Number(price),
                }),
            });

            if (!res.ok) throw new Error('Gagal simpan data');

            const json = await res.json();
            onAddMedicine(json.data);

            setName('');
            setCategory('');
            setStock('');
            setUnit('');
            setPrice('');
        } catch (err) {
            alert('Error: ' + err.message);
        } finally {
            setSubmitting(false);
        }
    };

    return (
        <form
            onSubmit={handleSubmit}
            className="bg-white rounded-xl shadow-sm border border-slate-200 p-6 mb-6"
        >
            <h2 className="text-lg font-semibold text-slate-800 mb-4">
                Tambah Obat
            </h2>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">
                        Nama Obat
                    </label>
                    <input
                        type="text"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        placeholder="Paracetamol 500mg"
                        className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>

                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">
                        Kategori
                    </label>
                    <input
                        type="text"
                        value={category}
                        onChange={(e) => setCategory(e.target.value)}
                        placeholder="Analgesik"
                        className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>

                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">
                        Stok
                    </label>
                    <input
                        type="number"
                        value={stock}
                        onChange={(e) => setStock(e.target.value)}
                        placeholder="100"
                        className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>

                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">
                        Satuan
                    </label>
                    <input
                        type="text"
                        value={unit}
                        onChange={(e) => setUnit(e.target.value)}
                        placeholder="tablet / botol / box"
                        className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>

                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">
                        Harga
                    </label>
                    <input
                        type="number"
                        value={price}
                        onChange={(e) => setPrice(e.target.value)}
                        placeholder="5000"
                        className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                </div>
            </div>

            <button
                type="submit"
                disabled={submitting}
                className="mt-4 bg-blue-600 hover:bg-blue-700 disabled:bg-blue-300 text-white text-sm font-medium px-4 py-2 rounded-lg transition-colors"
            >
                {submitting ? 'Menyimpan...' : 'Simpan Obat'}
            </button>
        </form>
    );
}

export default MedicineForm;