const fetch = require('node-fetch');

class PrayerTimesResponse {
  constructor(data) {
    this.CityName = data.CityName;
    this.CountryCode = data.CountryCode;
    this.CountryName = data.CountryName;
    this.CityLName = data.CityLName;
    this.CountryLName = data.CountryLName;
    this.CountryAlpha2 = data.CountryAlpha2;
    this.TimeZone = data.TimeZone;
    this.Imsaak = data.Imsaak;
    this.Sunrise = data.Sunrise;
    this.SunriseDT = data.SunriseDT;
    this.Noon = data.Noon;
    this.Sunset = data.Sunset;
    this.Maghreb = data.Maghreb;
    this.Midnight = data.Midnight;
    this.Today = data.Today;
    this.TodayQamari = data.TodayQamari;
    this.TodayGregorian = data.TodayGregorian;
    this.DayLength = data.DayLength;
    this.SimultaneityOfKaaba = data.SimultaneityOfKaaba;
    this.SimultaneityOfKaabaDesc = data.SimultaneityOfKaabaDesc;
  }
}
async function fetchData() {
  const url = "https://prayer.aviny.com/api/prayertimes/11";
  try {
    const response = await fetch(url);
    const data = await response.json();
    const result = new PrayerTimesResponse(data);

    const loc = 'Asia/Tehran';
    const now = new Date();
    const year = now.getFullYear();
    const month = now.getMonth() + 1;
    const day = now.getDate();
    
    const sobh = new Date(year, month - 1, day, ...result.Imsaak.split(':'));
    const zohr = new Date(year, month - 1, day, ...result.Noon.split(':'));
    const magh = new Date(year, month - 1, day, ...result.Maghreb.split(':'));
    
    let duration;
    if (now < sobh) {
      duration = sobh - now;
    } else if (now < zohr) {
      duration = zohr - now;
    } else if (now < magh) {
      duration = magh - now;
    } else {
      const nextSobh = new Date(year, month - 1, day + 1, ...result.Imsaak.split(':'));
      duration = nextSobh - now;
    }
    
    console.log(Math.floor(duration / 60000) + " minutes");
  } catch (error) {
    console.error("Error:", error);
  }
}

fetchData();
