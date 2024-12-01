/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.kaz.aoc.day1;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Scanner;
import java.math;

/**
 *
 * @author dhc10
 */
public class Day1 {
	
    public static void main(String args[]) {
        ArrayList<Integer> listOne = new ArrayList<Integer>();
        ArrayList<Integer> listTwo = new ArrayList<Integer>();
         try {
            File lists = new File("../inputFiles/Day1.txt");
            Scanner myReader = new Scanner(lists);
            while(myReader.hasNextLine()) {
                String bothList = myReader.nextLine();
                String[] splitLists = bothList.split("   ");
                listOne.add(Integer.decode(splitLists[0]));
                listTwo.add(Integer.decode(splitLists[1]));
            }
            myReader.close();
        } catch (FileNotFoundException e) {
            System.out.println("No File Found");
            e.printStackTrace();
        }
        listOne.sort(null);
        listTwo.sort(null);
        Integer totalDifference = 0;
        for (int i=0;i<listOne.size();i++) {
            totalDifference += Math.abs(listOne.get(i)-listTwo.get(i));            
        }
        Integer similarityScore = 0;
        System.out.printf("Part 1 Answer: %d\n", totalDifference);
        for (Integer i : listOne) {
            similarityScore += i*Collections.frequency(listTwo, i);
        }
        System.out.printf("Part 2 Answer: %d\n", similarityScore);
    }
}
