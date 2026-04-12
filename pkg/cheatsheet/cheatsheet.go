package cheatsheet

const Content = `
# TERMINAL
ls -la              → tüm dosyaları listele (gizliler dahil)
ls -lh              → okunabilir boyutla listele
cd ~                → home klasörüne git
pwd                 → bulunduğun yeri göster
mkdir -p a/b/c      → iç içe klasör oluştur
rm -rf dir          → klasörü sil (⚠️ dikkat)
cp -r src dst       → klasörü kopyala
mv src dst          → taşı / yeniden adlandır
find . -name "*.log"          → log dosyalarını bul
find . -size +100M            → 100MB+ dosyaları bul
du -sh * | sort -rh | head -5 → en büyük 5 klasör
df -h               → disk kullanımı
free -h             → RAM kullanımı
cat file            → dosya içeriğini göster
tail -f file        → canlı izle
grep -r "text" .    → klasörde metin ara
grep -i "err" app.log → hata satırlarını bul
ps aux              → çalışan processleri göster
kill -9 PID         → process sonlandır
chmod +x file       → çalıştırılabilir yap
curl -I url         → HTTP header'larını gör
wget url            → dosya indir
tar -czf out.tar.gz dir  → sıkıştır
tar -xzf file.tar.gz     → aç
history | grep git  → geçmişte ara
export VAR=val      → değişken tanımla
source ~/.bashrc    → config'i yenile
ssh user@host       → uzak sunucuya bağlan

# GIT
git init            → repo başlat
git clone url       → repoyu indir
git status          → değişiklikleri göster
git add .           → tümünü hazırla
git commit -m "msg" → commit yap
git push            → uzak repoya gönder
git pull            → uzak repodan çek
git log --oneline   → commit geçmişi
git log -5          → son 5 commit
git log --since="7 days ago" → son 7 günün commitleri
git diff            → değişiklikleri göster
git diff --name-only → sadece değişen dosyalar
git branch          → branch'leri listele
git checkout -b name → branch oluştur ve geç
git merge branch    → birleştir
git stash           → geçici sakla
git stash pop       → geri getir
git reset --hard HEAD → tüm değişiklikleri geri al (⚠️)
git reset HEAD~1    → son commiti geri al
git revert hash     → commiti geri döndür
git blame file      → satır satır kim değiştirdi
git shortlog -sn    → kişi başına commit sayısı

# DOCKER
docker ps           → çalışan container'lar
docker ps -a        → tüm container'lar
docker images       → image listesi
docker pull image   → image indir
docker run -d image → arka planda çalıştır
docker run -it image sh → interaktif shell
docker stop id      → durdur
docker rm id        → container sil
docker rmi image    → image sil
docker logs -f id   → logları canlı izle
docker exec -it id sh → çalışan container'a gir
docker build -t name . → image oluştur
docker-compose up -d   → servisleri başlat
docker-compose down    → servisleri durdur
docker system prune -f → temizle
docker stats           → kaynak kullanımı

# NPM / YARN
npm install         → bağımlılıkları yükle
npm run dev         → geliştirme sunucusunu başlat
npm run build       → projeyi derle
npm test            → testleri çalıştır
npm outdated        → eski paketleri listele
npm audit fix       → güvenlik açıklarını düzelt
yarn add package    → paket ekle
yarn remove package → paket kaldır
`

const SystemPrompt = `Sen pilot adlı bir terminal asistanısın. SADECE terminal/git/docker komutları üretirsin.

ZORUNLU KURALLAR:
1. Komut üretirken YALNIZCA şu formatı kullan, başka hiçbir şey yazma:
` + "```" + `
komut buraya
` + "```" + `
📌 Ne yapar: tek cümle Türkçe açıklama

2. Açıklama yaparken YALNIZCA şu formatı kullan:
🔍 Bu komut: tek cümle açıklama
📦 Parçalar:
  - parça1: açıklama
  - parça2: açıklama

3. Tehlikeli komutlar için ⚠️ ekle
4. Birden fazla işlem gerekiyorsa && veya | kullan
5. ASLA hikaye yazma, ASLA gereksiz açıklama yapma
6. Cheatsheet dışında komut bilmiyorsan "Bu komutu bilmiyorum." de

Cheatsheet:
` + Content