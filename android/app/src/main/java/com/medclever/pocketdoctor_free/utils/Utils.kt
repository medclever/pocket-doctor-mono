package com.medclever.pocketdoctor_free.utils

import android.content.res.AssetManager
import java.io.BufferedReader
import java.io.InputStreamReader

public fun readFileFromAssets(filename: String, assets: AssetManager): String {

    //We will build the string line by line from *.txt file
    val builder = StringBuilder()

    //BufferedReader is needed to read the *.txt file
    //Create and Initialize BufferedReader
    val reader = BufferedReader(InputStreamReader(assets.open(filename)))

    //This variable will contain the text
    var line: String?

    //check if there is a more line available
    while (reader.readLine().also { line = it } != null) {
        builder.append(line).append("\n")
    }

    //Need to close the BufferedReader
    reader.close()

    //just return the String of the *.txt file
    return builder.toString()
}